package main

import (
	"os"
	log "github.com/Sirupsen/logrus"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/aws/aws-sdk-go/aws"
	"encoding/json"
	"time"
	"github.com/xtracdev/pgpublish"
	"github.com/golang/protobuf/proto"
	"github.com/xtracdev/synthpub/synthevent"
)

const (
	QueueUrlEnv         = "EVENT_QUEUE_URL"
	MaxMessages = 50
)

type SNSMessage struct {
	Message string
}

func SNSMessageFromRawMessage(raw string) (*SNSMessage, error) {
	var snsMessage SNSMessage
	err := json.Unmarshal([]byte(raw), &snsMessage)
	return &snsMessage, err
}

func dumpMessage(msg string) {
	var aggId, typecode string
	var version int
	var payload []byte
	var err error
	var timestamp time.Time

	aggId, version, payload, typecode, timestamp, err = pgpublish.DecodePGEvent(msg)
	if err != nil {
		log.Infof("Error decoding message", err.Error())
		return
	}

	log.Infof("%s event %s %d %v", typecode, aggId, version, timestamp)

	if typecode != synthevent.EventTypeCode {
		log.Infof("Hmmm... I'm not interested in %s messages", typecode)
		return
	}

	var unpickled synthevent.SyntheticEvent
	err = proto.Unmarshal(payload, &unpickled)
	switch err {
	case nil:
		log.Infof("Injection time: %v", unpickled.GetInjectedTime())
	default:
		log.Warnf("Error unmarshalling payload: %v", err.Error())
	}
}

func main() {

	queueURL := os.Getenv(QueueUrlEnv)
	if queueURL == "" {
		log.Fatalf("Queue url not available via %s", QueueUrlEnv)
	}

	log.Info("Create session")
	session, err := session.NewSession()
	if err != nil {
		log.Fatal(err.Error())
	}

	svc := sqs.New(session)

	messageCount := 0

	//While there be messages to consume, consume messages...
	for {
		messageCount++
		if messageCount == MaxMessages {
			log.Infof("processed %d messages - exiting", messageCount)
			return
		}

		params := &sqs.ReceiveMessageInput{
			QueueUrl:            aws.String(queueURL), // Required
			MaxNumberOfMessages: aws.Int64(1),
			WaitTimeSeconds:     aws.Int64(10),
		}

		log.Debug("Receieve message")
		resp, err := svc.ReceiveMessage(params)
		if err != nil {
			log.Fatal(err.Error())
		}

		messages := resp.Messages
		if len(messages) == 0 {
			log.Info("No message available within timeout window - exiting.")
			return
		}

		message := *messages[0]
		log.Infof("Message: %v", message)

		sns, err := SNSMessageFromRawMessage(*message.Body)
		if err != nil {
			log.Fatal(err.Error())
		}

		dumpMessage(sns.Message)

		deleteParams := &sqs.DeleteMessageInput{
			QueueUrl:      aws.String(queueURL),
			ReceiptHandle: message.ReceiptHandle,
		}

		_, err = svc.DeleteMessage(deleteParams)
	}
}

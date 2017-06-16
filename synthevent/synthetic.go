package synthevent

import (
	"github.com/golang/protobuf/proto"
	"github.com/xtracdev/goes"
	"github.com/xtracdev/goes/uuid"
	"math/rand"
	"time"
)

const (
	EventTypeCode = "SYNTH"
)

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

func NewSyntheticEvent() (*SyntheticEvent, error) {
	id, err := uuid.GenerateUuidV4()
	if err != nil {
		return nil, err
	}

	now := time.Now()
	return &SyntheticEvent{
		EventId:  id,
		Injected: now.UnixNano(),
	}, nil
}

func (e *SyntheticEvent) GetInjectedTime() time.Time {
	return time.Unix(0, e.Injected)
}

func (e *SyntheticEvent) ToGoESEvent() (*goes.Event, error) {
	var newEvent goes.Event
	var err error
	newEvent.Source = e.EventId
	newEvent.Version = r.Int()
	newEvent.TypeCode = EventTypeCode
	newEvent.Payload, err = proto.Marshal(e)
	if err != nil {
		return nil, err
	}

	return &newEvent, nil
}

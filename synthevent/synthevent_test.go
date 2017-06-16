package synthevent

import (
	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshalling(t *testing.T) {

	synthEvent, err := NewSyntheticEvent()
	if assert.Nil(t, err) {
		marshalled, err := proto.Marshal(synthEvent)
		if assert.Nil(t, err) {
			var unpickled SyntheticEvent
			err := proto.Unmarshal(marshalled, &unpickled)
			if assert.Nil(t, err) {
				assert.Equal(t, synthEvent.EventId, unpickled.EventId)
				assert.True(t, synthEvent.GetInjectedTime().Equal(unpickled.GetInjectedTime()))
			}
		}
	}
}

func TestEventConversion(t *testing.T) {
	synthEvent, err := NewSyntheticEvent()
	if assert.Nil(t, err) {
		e, err := synthEvent.ToGoESEvent()
		if assert.Nil(t, err) {
			assert.Equal(t, synthEvent.EventId, e.Source)
			assert.Equal(t, EventTypeCode, e.TypeCode)

			var unpickled SyntheticEvent
			err = proto.Unmarshal(e.Payload.([]byte), &unpickled)
			if assert.Nil(t, err) {
				assert.Equal(t, synthEvent.EventId, unpickled.EventId)
				assert.True(t, synthEvent.GetInjectedTime().Equal(unpickled.GetInjectedTime()))
			}
		}
	}
}

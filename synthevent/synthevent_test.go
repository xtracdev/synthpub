package synthevent

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/golang/protobuf/proto"
)

func TestMarshalling(t *testing.T) {

	synthEvent, err := NewSyntheticEvent()
	if assert.Nil(t,err) {
		marshalled,err := proto.Marshal(synthEvent)
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

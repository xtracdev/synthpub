package synthevent

import (
	"testing"
	"time"
	"github.com/xtracdev/goes/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/golang/protobuf/proto"
)

func TestMarshalling(t *testing.T) {
	id,err := uuid.GenerateUuidV4()
	if assert.Nil(t, err) {
		now := time.Now()
		synthEvent := SyntheticEvent{
			EventId:id,
			Injected:now.UnixNano(),
		}

		marshalled,err := proto.Marshal(&synthEvent)
		if assert.Nil(t, err) {
			var unpickled SyntheticEvent
			err := proto.Unmarshal(marshalled, &unpickled)
			if assert.Nil(t,err) {
				assert.Equal(t, id, unpickled.EventId)
				injectedNow := time.Unix(0, unpickled.Injected)
				assert.True(t, injectedNow.Equal(now))
			}
		}
	}
}

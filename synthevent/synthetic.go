package synthevent

import (
	"github.com/xtracdev/goes/uuid"
	"time"
)

func NewSyntheticEvent() (*SyntheticEvent,error) {
	id,err := uuid.GenerateUuidV4()
	if err != nil {
		return nil,err
	}

	now := time.Now()
	return &SyntheticEvent{
		EventId:id,
		Injected:now.UnixNano(),
	},nil
}

func (e *SyntheticEvent) GetInjectedTime() time.Time {
	return time.Unix(0, e.Injected)
}


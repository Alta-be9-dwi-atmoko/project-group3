package business

import (
	"project/group3/features/events"
)

type eventUseCase struct {
	eventData events.Data
}

func NewEventBusiness(evtData events.Data) events.Business {
	return &eventUseCase{
		eventData: evtData,
	}
}

func (uc *eventUseCase) CreateEventBusiness(newData events.Core) (response int, err error) {
	response, err = uc.eventData.InsertData(newData)

	return response, err
}

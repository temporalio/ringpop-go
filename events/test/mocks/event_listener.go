package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/temporalio/ringpop-go/events"
)

type EventListener struct {
	mock.Mock
}

// HandleEvent provides a mock function with given fields: event
func (_m *EventListener) HandleEvent(event events.Event) {
	_m.Called(event)
}

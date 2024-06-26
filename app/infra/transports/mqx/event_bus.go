package mqx

import (
	"github.com/blackhorseya/godine/entity/events"
)

// EventBus defines the interface for an event bus.
type EventBus interface {
	Register(eventType string, handler func(events.DomainEvent))
	Unregister(eventType string, handler func(events.DomainEvent))
	Publish(event events.DomainEvent)
}

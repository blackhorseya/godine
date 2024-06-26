package mqx

import (
	"github.com/blackhorseya/godine/entity/events"
)

// HandlerID defines the type for handler ID.
type HandlerID int

// EventBus defines the interface for an event bus.
type EventBus interface {
	Register(eventType string, handler func(events.DomainEvent)) HandlerID
	Unregister(eventType string, id HandlerID)
	Publish(event events.DomainEvent)
}

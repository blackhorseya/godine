package events

import (
	"github.com/blackhorseya/godine/pkg/contextx"
)

// EventHandler defines the interface for event handlers.
type EventHandler interface {
	Handle(ctx contextx.Contextx, event *DomainEvent) error
}

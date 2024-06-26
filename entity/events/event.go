package events

import (
	"time"

	"github.com/blackhorseya/godine/pkg/contextx"
)

// DomainEvent is the interface for domain events.
type DomainEvent interface {
	OccurredOn(ctx contextx.Contextx) time.Time
	EventType() string
}

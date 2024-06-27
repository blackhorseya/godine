package mqx

import (
	"encoding/json"
	"sync"

	"github.com/blackhorseya/godine/app/infra/transports/kafkax"
	"github.com/blackhorseya/godine/entity/events"
	"github.com/blackhorseya/godine/pkg/contextx"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
)

const topic = "new_events"

type KafkaEventBus struct {
	reader   *kafka.Reader
	writer   *kafka.Writer
	handlers map[string]map[HandlerID]func(events.DomainEvent)
	nextID   HandlerID
	mu       sync.RWMutex
}

// NewKafkaEventBus creates a new Kafka event bus
func NewKafkaEventBus() (EventBus, error) {
	reader, err := kafkax.NewReaderWithTopic(topic)
	if err != nil {
		return nil, err
	}

	writer, err := kafkax.NewWriter()
	if err != nil {
		return nil, err
	}

	bus := &KafkaEventBus{
		reader:   reader,
		writer:   writer,
		handlers: make(map[string]map[HandlerID]func(events.DomainEvent)),
		nextID:   0,
	}

	go bus.startConsuming()

	return bus, nil
}

func (bus *KafkaEventBus) startConsuming() {
	for {
		ctx := contextx.Background()
		m, err := bus.reader.ReadMessage(ctx)
		if err != nil {
			ctx.Error("Error reading message", zap.Error(err))
			continue
		}

		var event events.DomainEvent
		if err = json.Unmarshal(m.Value, &event); err != nil {
			ctx.Error("Error unmarshalling event", zap.Error(err))
			continue
		}

		bus.mu.RLock()
		handlers, found := bus.handlers[event.Topic()]
		bus.mu.RUnlock()

		if found {
			for _, handler := range handlers {
				go handler(event)
			}
		}
	}
}

// Register registers an event handler and returns a unique handler ID
func (bus *KafkaEventBus) Register(eventType string, handler func(events.DomainEvent)) HandlerID {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	if _, ok := bus.handlers[eventType]; !ok {
		bus.handlers[eventType] = make(map[HandlerID]func(events.DomainEvent))
	}
	id := bus.nextID
	bus.handlers[eventType][id] = handler
	bus.nextID++
	return id
}

// Unregister unregisters an event handler by its ID
func (bus *KafkaEventBus) Unregister(eventType string, id HandlerID) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	if handlers, ok := bus.handlers[eventType]; ok {
		delete(handlers, id)
		if len(handlers) == 0 {
			delete(bus.handlers, eventType)
		}
	}
}

// Publish publishes an event to all registered handlers
func (bus *KafkaEventBus) Publish(ctx contextx.Contextx, event events.DomainEvent) {
	data, err := json.Marshal(event)
	if err != nil {
		ctx.Error("Error marshalling event", zap.Error(err), zap.Any("event", event))
		return
	}

	msg := kafka.Message{
		Topic: event.Topic(),
		Value: data,
	}

	if err = bus.writer.WriteMessages(ctx, msg); err != nil {
		ctx.Error("Error writing message", zap.Error(err))
	}
}

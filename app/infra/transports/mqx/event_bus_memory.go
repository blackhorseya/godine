package mqx

import (
	"sync"

	"github.com/blackhorseya/godine/entity/events"
	"github.com/blackhorseya/godine/pkg/contextx"
)

type MemoryEventBus struct {
	handlers map[string]map[HandlerID]func(events.DomainEvent)
	nextID   HandlerID
	mu       sync.RWMutex
}

// NewMemoryEventBus 創建一個新的內存事件總線
func NewMemoryEventBus() EventBus {
	return &MemoryEventBus{
		handlers: make(map[string]map[HandlerID]func(events.DomainEvent)),
		nextID:   0,
	}
}

// Register 註冊事件處理器並返回處理器的唯一標識符
func (bus *MemoryEventBus) Register(eventType string, handler func(events.DomainEvent)) HandlerID {
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

// Unregister 取消註冊事件處理器
func (bus *MemoryEventBus) Unregister(eventType string, id HandlerID) {
	bus.mu.Lock()
	defer bus.mu.Unlock()
	if handlers, ok := bus.handlers[eventType]; ok {
		delete(handlers, id)
		if len(handlers) == 0 {
			delete(bus.handlers, eventType)
		}
	}
}

// Publish 發布事件
func (bus *MemoryEventBus) Publish(ctx contextx.Contextx, event events.DomainEvent) {
	bus.mu.RLock()
	defer bus.mu.RUnlock()
	if handlers, found := bus.handlers[event.EventType()]; found {
		for _, handler := range handlers {
			handler(event)
		}
	}
}

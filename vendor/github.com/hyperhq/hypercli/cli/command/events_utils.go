package command

import (
	"sync"

	eventtypes "github.com/hyperhq/hyper-api/types/events"
)

type eventProcessor func(eventtypes.Message, error) error

// EventHandler is abstract interface for user to customize
// own handle functions of each type of events
type EventHandler interface {
	Handle(action string, h func(eventtypes.Message))
	Watch(c <-chan eventtypes.Message)
}

// InitEventHandler initializes and returns an EventHandler
func InitEventHandler() EventHandler {
	return &eventHandler{handlers: make(map[string]func(eventtypes.Message))}
}

type eventHandler struct {
	handlers map[string]func(eventtypes.Message)
	mu       sync.Mutex
}

func (w *eventHandler) Handle(action string, h func(eventtypes.Message)) {
	w.mu.Lock()
	w.handlers[action] = h
	w.mu.Unlock()
}

// Watch ranges over the passed in event chan and processes the events based on the
// handlers created for a given action.
// To stop watching, close the event chan.
func (w *eventHandler) Watch(c <-chan eventtypes.Message) {
	for e := range c {
		w.mu.Lock()
		h, exists := w.handlers[e.Action]
		w.mu.Unlock()
		if !exists {
			continue
		}
		go h(e)
	}
}

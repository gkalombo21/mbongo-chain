package network

import (
	"context"
	"sync"
)

// Router offers a minimal publish-subscribe mechanism for runtime components.
type Router struct {
	lock        sync.RWMutex
	subscribers map[string][]chan []byte
}

// NewRouter instantiates an empty router.
func NewRouter() *Router {
	return &Router{
		subscribers: make(map[string][]chan []byte),
	}
}

// Subscribe registers a channel to receive messages for a topic.
func (r *Router) Subscribe(topic string, ch chan []byte) {
	r.lock.Lock()
	defer r.lock.Unlock()

	r.subscribers[topic] = append(r.subscribers[topic], ch)
}

// Publish distributes data to topic subscribers.
// TODO: add buffering, delivery guarantees, and metrics hooks.
func (r *Router) Publish(_ context.Context, topic string, payload []byte) {
	r.lock.RLock()
	defer r.lock.RUnlock()

	for _, ch := range r.subscribers[topic] {
		select {
		case ch <- payload:
		default:
			// Drop message if subscriber is not ready. Future versions may queue.
		}
	}
}

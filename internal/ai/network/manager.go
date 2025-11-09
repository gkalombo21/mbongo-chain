package network

import (
	"fmt"
	"sync"
	"time"
)

// Manager controls message flow between network participants.
type Manager struct {
	network *Network

	inbox chan Message
	quit  chan struct{}

	mu sync.RWMutex
}

// NewManager creates a manager with buffered message channel.
func NewManager(net *Network) *Manager {
	if net == nil {
		net = NewNetwork()
	}

	return &Manager{
		network: net,
		inbox:   make(chan Message, 64),
		quit:    make(chan struct{}),
	}
}

// Start begins background processing of incoming messages.
func (m *Manager) Start() {
	fmt.Println("🛰️  AI Network manager online")
	go m.loop()
}

// Stop halts the manager.
func (m *Manager) Stop() {
	close(m.quit)
}

// BroadcastJobRequest sends a job to all known nodes (simulated).
func (m *Manager) BroadcastJobRequest(req JobRequest) {
	m.mu.RLock()
	nodes := m.network.ListNodes()
	m.mu.RUnlock()

	for _, node := range nodes {
		fmt.Printf("📡 Sending job %s to node %s@%s\n", req.JobID, node.ID, node.Address)
	}
}

// SubmitResult pushes a job result into the manager's inbox.
func (m *Manager) SubmitResult(result JobResult) {
	select {
	case m.inbox <- result:
	default:
		fmt.Println("⚠️  network inbox full; dropping result", result.JobID)
	}
}

func (m *Manager) loop() {
	for {
		select {
		case msg := <-m.inbox:
			m.handleMessage(msg)
		case <-time.After(30 * time.Second):
			fmt.Println("🔁 Network manager idle heartbeat")
		case <-m.quit:
			fmt.Println("🛑 Network manager stopped")
			return
		}
	}
}

func (m *Manager) handleMessage(msg Message) {
	switch v := msg.(type) {
	case JobResult:
		fmt.Printf("📥 Received result for job %s from %s (accuracy %.2f)\n", v.JobID, v.NodeID, v.Accuracy)
	default:
		fmt.Println("ℹ️  Received message", v)
	}
}

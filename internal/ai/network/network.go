package network

import (
	"fmt"
	"sync"
)

// Network coordinates peer-to-peer style messaging between compute nodes.
type Network struct {
	mu    sync.RWMutex
	nodes map[string]*Node
}

// NewNetwork constructs an empty network registry.
func NewNetwork() *Network {
	return &Network{nodes: make(map[string]*Node)}
}

// RegisterNode adds or updates a node in the network registry.
func (n *Network) RegisterNode(node *Node) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.nodes[node.ID] = node
	fmt.Printf("🌐 Registered compute node %s (%s)\n", node.ID, node.Type)
}

// ListNodes returns a snapshot of known nodes.
func (n *Network) ListNodes() []*Node {
	n.mu.RLock()
	defer n.mu.RUnlock()

	list := make([]*Node, 0, len(n.nodes))
	for _, node := range n.nodes {
		list = append(list, node)
	}
	return list
}

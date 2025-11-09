package network

import "time"

// Node describes an AI compute participant known to the network manager.
type Node struct {
	ID         string
	Address    string
	Type       string // gpu or cpu
	ActiveJobs []string
	LastSeen   time.Time
}

// UpdateHeartbeat refreshes bookkeeping details when a heartbeat is received.
func (n *Node) UpdateHeartbeat(h Heartbeat) {
	n.Address = h.Address
	n.ActiveJobs = make([]string, h.ActiveJobs)
	n.LastSeen = time.Now()
}

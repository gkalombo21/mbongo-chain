package network

// Message is the empty marker interface implemented by all network messages.
type Message interface{}

// JobRequest is broadcast to available compute nodes when work is available.
type JobRequest struct {
	JobID      string
	Model      string
	PayloadURI string
	RewardHint float64
	Priority   int
}

// JobResult is submitted by compute nodes once a task is complete.
type JobResult struct {
	JobID      string
	NodeID     string
	OutputHash string
	Accuracy   float64
	RuntimeMs  int64
}

// Heartbeat is periodically exchanged between nodes to confirm liveness.
type Heartbeat struct {
	NodeID     string
	Address    string
	ActiveJobs int
}

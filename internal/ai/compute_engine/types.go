package computeengine

// Job represents a single AI workload dispatched to the PoUW network.
type Job struct {
	ID            string
	Model         string
	DataURI       string
	RewardPerUnit float64
	MaxRuntimeMs  int64
	Metadata      map[string]string
}

// ComputeNode represents a worker node in the network.
type ComputeNode struct {
	ID     string
	Type   string // gpu or cpu
	Status string
}

// NodeInfo captures runtime details about the node executing a job.
type NodeInfo struct {
	ID         string
	Type       string // gpu or cpu
	Region     string
	Reputation float64
}

// Result contains the output metrics produced by a compute node.
type Result struct {
	JobID       string
	NodeID      string
	OutputHash  string
	Accuracy    float64
	RuntimeMs   int64
	ProofString string
	Metadata    map[string]any
}

// Proof encapsulates the cryptographic information submitted by the node.
type Proof struct {
	JobID    string
	NodeID   string
	ProofHex string
	Hash     string
	Accuracy float64
}

// RewardDecision captures information required to distribute MBG rewards.
type RewardDecision struct {
	JobID           string
	NodeID          string
	RewardAmountMBG float64
	Reason          string
}

package computeengine

// WorkRequest captures metadata required to execute a useful-work task.
// Additional scheduling hints and security attributes will be added later.
type WorkRequest struct {
	ID          string
	Model       string
	Submitter   string
	InputURI    string
	MaxDuration int64
}

// WorkResult represents the output produced by a compute operator.
type WorkResult struct {
	RequestID string
	OutputURI string
	ProofID   string
}

// Scheduler defines the interface for enqueueing and dispatching work.
type Scheduler interface {
	Submit(request WorkRequest) error
	Next() (WorkRequest, bool)
}

// Executor defines the interface for executing work requests.
type Executor interface {
	Execute(request WorkRequest) (WorkResult, error)
}

// TODO: wire Scheduler implementations to job queues and integrate runtime.

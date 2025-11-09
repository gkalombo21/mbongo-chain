package computeengine

import (
	"errors"
	"fmt"
)

// RewardCalculator encapsulates how MBG payouts are computed for a job result.
type RewardCalculator interface {
	Calculate(result Result) (RewardDecision, error)
}

// RewardFunc allows simple functions to satisfy RewardCalculator.
type RewardFunc func(Result) (RewardDecision, error)

// Calculate implements RewardCalculator.
func (f RewardFunc) Calculate(result Result) (RewardDecision, error) {
	return f(result)
}

// Engine orchestrates job scheduling, proof verification, and reward calculation.
type Engine struct {
	queue      JobQueue
	verifier   ProofVerifier
	rewardCalc RewardCalculator
}

// NewEngine constructs a new Engine with sensible defaults.
func NewEngine(queue JobQueue, verifier ProofVerifier, calculator RewardCalculator) *Engine {
	if queue == nil {
		queue = NewMemoryQueue(0)
	}

	if verifier == nil {
		verifier = NewSimpleVerifier()
	}

	if calculator == nil {
		calculator = RewardFunc(defaultRewardCalculation)
	}

	return &Engine{
		queue:      queue,
		verifier:   verifier,
		rewardCalc: calculator,
	}
}

// ScheduleJob enqueues a job for later dispatch to compute nodes.
func (e *Engine) ScheduleJob(job Job) error {
	fmt.Println("enqueue job", job.ID)
	return e.queue.Enqueue(job)
}

// NextJob retrieves the next available job for dispatch.
func (e *Engine) NextJob() (Job, error) {
	job, err := e.queue.Dequeue()
	if err != nil {
		return Job{}, err
	}

	fmt.Println("dispatch job", job.ID)
	return job, nil
}

// HandleResult processes the proof and computes the reward decision.
func (e *Engine) HandleResult(result Result, proof Proof) (RewardDecision, error) {
	if result.JobID == "" || proof.JobID == "" {
		return RewardDecision{}, errors.New("job id required for verification")
	}

	if result.JobID != proof.JobID {
		return RewardDecision{}, errors.New("result and proof job ids do not match")
	}

	if err := e.verifier.Verify(proof); err != nil {
		return RewardDecision{}, err
	}

	reward, err := e.rewardCalc.Calculate(result)
	if err != nil {
		return RewardDecision{}, err
	}

	fmt.Println("reward decision", reward.JobID, reward.RewardAmountMBG)
	// TODO: Integrate reward settlement with Cosmos SDK bank module.
	return reward, nil
}

// defaultRewardCalculation applies a simple formula compatible with the AI policy guide.
func defaultRewardCalculation(result Result) (RewardDecision, error) {
	if result.Accuracy < 0 {
		return RewardDecision{}, errors.New("accuracy cannot be negative")
	}

	reward := result.Accuracy * 0.1
	return RewardDecision{
		JobID:           result.JobID,
		NodeID:          result.NodeID,
		RewardAmountMBG: reward,
		Reason:          "simple reward function",
	}, nil
}

package compute_engine

import (
	"context"
	"errors"

	ai "github.com/gkalombo21/mbongo-chain/internal/ai/compute_engine"
)

// Pipeline coordinates the lifecycle of Proof-of-Useful-Work tasks by
// delegating to the existing scheduler, executor, and scoring components.
type Pipeline struct {
	Scheduler ai.Scheduler
	Executor  ai.Executor
	Scorer    ai.Scorer
}

// NewPipeline assembles a pipeline with the provided dependencies.
func NewPipeline(scheduler ai.Scheduler, executor ai.Executor, scorer ai.Scorer) *Pipeline {
	return &Pipeline{
		Scheduler: scheduler,
		Executor:  executor,
		Scorer:    scorer,
	}
}

// Dispatch enqueues work for future processing.
// TODO: add admission control and staking requirements.
func (p *Pipeline) Dispatch(_ context.Context, request ai.WorkRequest) error {
	if p.Scheduler == nil {
		return errors.New("scheduler not configured")
	}
	return p.Scheduler.Submit(request)
}

// ProcessNext pulls the next request from the scheduler and executes it.
// TODO: integrate operator identity and result notarization.
func (p *Pipeline) ProcessNext(_ context.Context) (*ai.WorkResult, error) {
	if p.Scheduler == nil || p.Executor == nil {
		return nil, errors.New("pipeline not fully configured")
	}

	request, ok := p.Scheduler.Next()
	if !ok {
		return nil, nil
	}

	result, err := p.Executor.Execute(request)
	if err != nil {
		if p.Scorer != nil {
			_ = p.Scorer.RecordFailure(request.Submitter, request, err)
		}
		return nil, err
	}

	if p.Scorer != nil {
		_ = p.Scorer.RecordSuccess(request.Submitter, result)
	}

	return &result, nil
}

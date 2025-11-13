package runtime

import (
	"context"

	"github.com/gkalombo21/mbongo-chain/internal/blockchain"
	internalcompute "github.com/gkalombo21/mbongo-chain/internal/compute_engine"
)

// Engine coordinates block execution, state management, and useful-work
// pipelines. It represents the smallest runnable core for the validator node.
type Engine struct {
	state     *blockchain.State
	executor  BlockExecutor
	pipeline  *internalcompute.Pipeline
	validator blockchain.Validator
}

// NewEngine builds a runtime engine with the provided dependencies. The state
// pointer is reused across executions, enabling callers to inspect the latest
// ledger snapshot after each block.
func NewEngine(state *blockchain.State, executor BlockExecutor, pipeline *internalcompute.Pipeline, validator blockchain.Validator) *Engine {
	return &Engine{
		state:     state,
		executor:  executor,
		pipeline:  pipeline,
		validator: validator,
	}
}

// Init prepares the engine to process blocks. Future implementations may load
// snapshots or hydrate caches.
func (e *Engine) Init(ctx context.Context) error {
	_ = ctx
	// TODO: hydrate state from persistent storage and warm caches.
	return nil
}

// ApplyBlock runs the transition pipeline for a single block.
func (e *Engine) ApplyBlock(ctx context.Context, block blockchain.Block) error {
	_ = ctx
	if e.executor == nil {
		return nil
	}
	return e.executor.Execute(block)
}

// Pipeline exposes the underlying useful-work pipeline.
func (e *Engine) Pipeline() *internalcompute.Pipeline {
	return e.pipeline
}

// State returns the current working state pointer for external inspection.
func (e *Engine) State() *blockchain.State {
	return e.state
}

// Validator returns the validator instance used by this engine.
func (e *Engine) Validator() blockchain.Validator {
	return e.validator
}

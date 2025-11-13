package runtime

import (
	"github.com/gkalombo21/mbongo-chain/internal/blockchain"
)

// BlockExecutor defines the behavior expected from a component capable of
// applying a block to the working state.
type BlockExecutor interface {
	Execute(block blockchain.Block) error
}

// DefaultBlockExecutor combines a blockchain state pointer with a transition
// pipeline to process blocks sequentially.
type DefaultBlockExecutor struct {
	State      *blockchain.State
	Transition blockchain.TransitionPipeline
}

// NewDefaultBlockExecutor returns an executor ready for use by the runtime.
func NewDefaultBlockExecutor(state *blockchain.State, pipeline blockchain.TransitionPipeline) *DefaultBlockExecutor {
	return &DefaultBlockExecutor{
		State:      state,
		Transition: pipeline,
	}
}

// Execute triggers the transition pipeline against the maintained state.
func (e *DefaultBlockExecutor) Execute(block blockchain.Block) error {
	if e.State == nil {
		return nil
	}
	return e.Transition.Execute(e.State, block)
}

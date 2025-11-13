package blockchain

// TransitionContext bundles together stateful objects used while processing a
// block. Runtime components can extend this structure as the protocol evolves.
type TransitionContext struct {
	State *State
	Block Block
}

// TransitionHandler defines hooks that allow modules to react before, during,
// and after block execution. Each hook is optional and can be left as a no-op.
type TransitionHandler interface {
	BeginBlock(ctx TransitionContext) error
	DeliverTx(ctx TransitionContext, tx Transaction) error
	EndBlock(ctx TransitionContext) error
}

// TransitionPipeline orchestrates block validation and state transitions.
type TransitionPipeline struct {
	Validator Validator
	Handler   TransitionHandler
}

// NewTransitionPipeline creates a pipeline with optional validator and handler.
func NewTransitionPipeline(validator Validator, handler TransitionHandler) TransitionPipeline {
	return TransitionPipeline{
		Validator: validator,
		Handler:   handler,
	}
}

// Execute applies a block to state, invoking handlers along the way.
// TODO: extend with event emission and proof verification.
func (p TransitionPipeline) Execute(state *State, block Block) error {
	if state == nil {
		return nil
	}

	if p.Validator != nil {
		if err := p.Validator.ValidateBlock(block); err != nil {
			return err
		}
	}

	ctx := TransitionContext{
		State: state,
		Block: block,
	}

	if p.Handler != nil {
		if err := p.Handler.BeginBlock(ctx); err != nil {
			return err
		}
	}

	for _, tx := range block.Transactions {
		if p.Validator != nil {
			if err := p.Validator.ValidateTransaction(tx); err != nil {
				return err
			}
		}

		if err := state.ApplyTransaction(tx); err != nil {
			return err
		}

		if p.Handler != nil {
			if err := p.Handler.DeliverTx(ctx, tx); err != nil {
				return err
			}
		}
	}

	if err := state.ApplyBlock(block); err != nil {
		return err
	}

	if p.Handler != nil {
		if err := p.Handler.EndBlock(ctx); err != nil {
			return err
		}
	}

	return nil
}

// NoopTransitionHandler is a helper that leaves all hooks empty.
type NoopTransitionHandler struct{}

func (NoopTransitionHandler) BeginBlock(TransitionContext) error { return nil }

func (NoopTransitionHandler) DeliverTx(TransitionContext, Transaction) error { return nil }

func (NoopTransitionHandler) EndBlock(TransitionContext) error { return nil }

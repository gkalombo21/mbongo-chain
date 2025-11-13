package blockchain

// State represents the chain state snapshot used while processing blocks.
// For MVP purposes this is held entirely in memory.
type State struct {
	Balances map[string]uint64
	Height   uint64
}

// NewState initializes an empty state with default values.
func NewState() *State {
	return &State{
		Balances: make(map[string]uint64),
		Height:   0,
	}
}

// ApplyTransaction mutates the state according to a single transaction.
// TODO: handle nonce tracking, fee distribution, and failure modes.
func (s *State) ApplyTransaction(tx Transaction) error {
	return nil
}

// ApplyBlock mutates state based on an entire block.
// TODO: enforce transaction ordering, proof verification, and rollback logic.
func (s *State) ApplyBlock(block Block) error {
	return nil
}


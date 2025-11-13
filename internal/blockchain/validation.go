package blockchain

// Validator is satisfied by any component capable of evaluating blocks or
// transactions for correctness.
type Validator interface {
	ValidateBlock(block Block) error
	ValidateTransaction(tx Transaction) error
}

// BasicValidator provides a simple, extensible implementation scaffold.
type BasicValidator struct{}

// ValidateBlock performs structural checks on block data.
// TODO: verify header fields, signatures, and state transitions.
func (BasicValidator) ValidateBlock(block Block) error {
	return nil
}

// ValidateTransaction performs structural checks on a transaction.
// TODO: ensure balances, signature validity, and payload constraints.
func (BasicValidator) ValidateTransaction(tx Transaction) error {
	return nil
}


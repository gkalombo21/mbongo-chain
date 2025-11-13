package blockchain

// Transaction represents the atomic unit processed by the blockchain.
// Fields are placeholders and will expand as transaction semantics mature.
type Transaction struct {
	ID        string
	Sender    string
	Recipient string
	Amount    uint64
	Payload   []byte
}

// TxEncoder defines the contract for encoding transactions into bytes.
type TxEncoder interface {
	Encode(tx Transaction) ([]byte, error)
}

// TxHasher defines the contract for hashing transactions.
type TxHasher interface {
	Hash(tx Transaction) (string, error)
}

// NewTransaction builds a transaction with the basics required for MVP flows.
// TODO: insert validation checks for payload size, signatures, and fees.
func NewTransaction(id, sender, recipient string, amount uint64, payload []byte) Transaction {
	return Transaction{
		ID:        id,
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Payload:   payload,
	}
}


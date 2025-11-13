package blockchain

// BlockHeader contains metadata required to identify a block and link it to
// the chain. The structure intentionally remains minimal at MVP stage.
type BlockHeader struct {
	Height    uint64
	ParentHash string
	Time       int64
	StateRoot  string
}

// Block represents the container for a collection of transactions and the
// corresponding header. Additional fields (e.g., signatures) will be added
// once consensus integration begins.
type Block struct {
	Header       BlockHeader
	Transactions []Transaction
}

// NewBlock builds a block instance from header data and a set of transactions.
// TODO: enforce validation and merkle root computation before construction.
func NewBlock(header BlockHeader, txs []Transaction) Block {
	return Block{
		Header:       header,
		Transactions: txs,
	}
}

// Hash computes the identifier for the block.
// TODO: implement cryptographic hashing and caching.
func (b Block) Hash() string {
	return ""
}


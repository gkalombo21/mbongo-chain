package bank

import (
	"sync"
	"time"
)

var (
	transactionsMu sync.RWMutex
	transactions   []Transaction
	nextTxID       = 1
)

// recordTransaction appends a new transaction to the in-memory history.
func recordTransaction(txType string, amount float64) {
	transactionsMu.Lock()
	tx := Transaction{
		ID:     nextTxID,
		Type:   txType,
		Amount: amount,
		Date:   time.Now(),
	}
	nextTxID++
	transactions = append(transactions, tx)
	transactionsMu.Unlock()
}

// GetTransactions returns a shallow copy of the current transaction history.
func GetTransactions() []Transaction {
	transactionsMu.RLock()
	defer transactionsMu.RUnlock()
	copySlice := make([]Transaction, len(transactions))
	copy(copySlice, transactions)
	return copySlice
}



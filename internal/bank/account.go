package bank

import (
	"fmt"
	"sync"
	"time"
)

var (
	accountsMu   sync.RWMutex
	accountsByID = make(map[int]*Account)
	nextAccountID = 1
)

// CreateAccount initializes a new bank account with the specified owner name and initial balance.
// Returns a pointer to the created Account with an auto-incremented ID and records the creation transaction.
func CreateAccount(owner string, initialBalance float64) *Account {
	if owner == "" {
		owner = "unknown"
	}
	if initialBalance < 0 {
		initialBalance = 0
	}

	accountsMu.Lock()
	id := nextAccountID
	nextAccountID++
	acct := &Account{
		ID:        id,
		OwnerName: owner,
		Balance:   initialBalance,
		CreatedAt: time.Now(),
	}
	accountsByID[id] = acct
	accountsMu.Unlock()

	recordTransaction("create", initialBalance)
	return acct
}

// Deposit adds the specified amount to the account balance.
// Amount must be positive.
func (a *Account) Deposit(amount float64) {
	if amount <= 0 {
		return
	}
	accountsMu.Lock()
	a.Balance += amount
	accountsMu.Unlock()
	recordTransaction("deposit", amount)
}

// Withdraw removes the specified amount from the account balance if sufficient funds exist.
// It returns an error if the amount is not positive or funds are insufficient.
func (a *Account) Withdraw(amount float64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be positive")
	}
	accountsMu.Lock()
	defer accountsMu.Unlock()
	if a.Balance < amount {
		return fmt.Errorf("insufficient funds: balance=%.2f want=%.2f", a.Balance, amount)
	}
	a.Balance -= amount
	recordTransaction("withdraw", amount)
	return nil
}

// GetBalance returns the current account balance.
func (a *Account) GetBalance() float64 {
	accountsMu.RLock()
	b := a.Balance
	accountsMu.RUnlock()
	return b
}

// GetAccount returns the account by ID if present.
func GetAccount(id int) (*Account, bool) {
	accountsMu.RLock()
	acct, ok := accountsByID[id]
	accountsMu.RUnlock()
	return acct, ok
}



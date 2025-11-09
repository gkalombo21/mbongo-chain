package bank

import (
	"fmt"
	"sync"
)

// AccountManager provides a simple, concurrency-safe in-memory store
// for bank accounts and supports create, deposit, withdraw and balance queries.
// This is intentionally minimal and in-memory for demonstration purposes.
type AccountManager struct {
	mu       sync.RWMutex
	balances map[string]int64
}

// NewAccountManager creates a new empty manager.
func NewAccountManager() *AccountManager {
	return &AccountManager{balances: make(map[string]int64)}
}

// CreateAccount initializes a new account in the AccountManager with the specified account ID.
// The account is created with zero balance. Returns an error if the account ID is empty or the account already exists.
func (m *AccountManager) CreateAccount(accountId string) error {
	if accountId == "" {
		return fmt.Errorf("account id must not be empty")
	}
	
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.balances[accountId]; exists {
		return fmt.Errorf("account %s already exists", accountId)
	}
	m.balances[accountId] = 0
	return nil
}

// Deposit adds a positive amount to the given account.
func (m *AccountManager) Deposit(accountId string, amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("deposit amount must be positive")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	if _, exists := m.balances[accountId]; !exists {
		return fmt.Errorf("account %s not found", accountId)
	}
	m.balances[accountId] += amount
	return nil
}

// Withdraw subtracts a positive amount from the given account if sufficient funds exist.
func (m *AccountManager) Withdraw(accountId string, amount int64) error {
	if amount <= 0 {
		return fmt.Errorf("withdraw amount must be positive")
	}
	m.mu.Lock()
	defer m.mu.Unlock()
	balance, exists := m.balances[accountId]
	if !exists {
		return fmt.Errorf("account %s not found", accountId)
	}
	if balance < amount {
		return fmt.Errorf("insufficient funds: balance=%d, want=%d", balance, amount)
	}
	m.balances[accountId] = balance - amount
	return nil
}

// Balance returns the current balance for an account.
func (m *AccountManager) Balance(accountId string) (int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	balance, exists := m.balances[accountId]
	if !exists {
		return 0, fmt.Errorf("account %s not found", accountId)
	}
	return balance, nil
}

// DefaultManager is a package-level manager for convenience in examples and simple CLIs.
var DefaultManager = NewAccountManager()

// CreateAccountByID creates a new account using the default AccountManager with the specified account ID.
// The account is initialized with zero balance. Returns an error if the account ID is empty or already exists.
func CreateAccountByID(accountId string) error { return DefaultManager.CreateAccount(accountId) }

// Deposit is a helper that delegates to DefaultManager.Deposit.
func Deposit(accountId string, amount int64) error { return DefaultManager.Deposit(accountId, amount) }

// Withdraw is a helper that delegates to DefaultManager.Withdraw.
func Withdraw(accountId string, amount int64) error { return DefaultManager.Withdraw(accountId, amount) }

// GetBalance is a helper that delegates to DefaultManager.Balance.
func GetBalance(accountId string) (int64, error) { return DefaultManager.Balance(accountId) }



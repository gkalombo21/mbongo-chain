package bank

import "time"

// Account represents a simple bank account with owner and balance.
type Account struct {
	ID        int
	OwnerName string
	Balance   float64
	CreatedAt time.Time
}

// Transaction represents a simple immutable record of an action on an account.
type Transaction struct {
	ID     int
	Type   string
	Amount float64
	Date   time.Time
}



package user

import "time"

// User represents an application user in Mbongo-Chain.
// It contains identity fields and creation timestamp.
type User struct {
	ID        int
	Name      string
	Email     string
	CreatedAt time.Time
}



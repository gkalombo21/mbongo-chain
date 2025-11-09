package user

import (
	"fmt"
	"regexp"
	"sync"
	"time"
)

var (
	usersMu    sync.RWMutex
	usersByID  = make(map[int]*User)
	nextUserID = 1
)

// emailRegex is a simple RFC5322-inspired pattern for demonstration. It is not fully strict.
var emailRegex = regexp.MustCompile(`^[A-Za-z0-9._%+-]+@[A-Za-z0-9.-]+\.[A-Za-z]{2,}$`)

// validateEmail returns an error if the email does not match a basic pattern.
func validateEmail(email string) error {
	if !emailRegex.MatchString(email) {
		return fmt.Errorf("invalid email format: %s", email)
	}
	return nil
}

// CreateUser creates and stores a new user with the provided name and email.
// The ID is auto-incremented and CreatedAt is set to the current time.
func CreateUser(name, email string) *User {
	if name == "" {
		name = "unknown"
	}
	if err := validateEmail(email); err != nil {
		// If invalid, keep empty to reflect a not yet valid email
		email = ""
	}

	usersMu.Lock()
	id := nextUserID
	nextUserID++
	u := &User{
		ID:        id,
		Name:      name,
		Email:     email,
		CreatedAt: time.Now(),
	}
	usersByID[id] = u
	usersMu.Unlock()
	return u
}

// UpdateEmail validates and updates the email of this user. If validation fails, an error is returned.
func (u *User) UpdateEmail(newEmail string) error {
	if err := validateEmail(newEmail); err != nil {
		return err
	}
	usersMu.Lock()
	u.Email = newEmail
	usersMu.Unlock()
	return nil
}

// DisplayInfo returns a human-readable string describing the user.
func (u *User) DisplayInfo() string {
	usersMu.RLock()
	defer usersMu.RUnlock()
	return fmt.Sprintf("User #%d: %s <%s> (created %s)", u.ID, u.Name, u.Email, u.CreatedAt.Format(time.RFC3339))
}

// GetUser retrieves a user by ID.
func GetUser(id int) (*User, bool) {
	usersMu.RLock()
	defer usersMu.RUnlock()
	u, ok := usersByID[id]
	return u, ok
}



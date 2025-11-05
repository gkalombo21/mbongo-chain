package types

import "fmt"

// MsgSend defines a simple transfer message.
type MsgSend struct {
    From   string
    To     string
    Amount int64
}

// ValidateBasic performs basic validation.
func (msg MsgSend) ValidateBasic() error {
    if msg.From == "" || msg.To == "" {
        return fmt.Errorf("sender and recipient addresses must not be empty")
    }
    if msg.Amount <= 0 {
        return fmt.Errorf("amount must be positive")
    }
    return nil
}

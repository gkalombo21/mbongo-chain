package bank

// Public interface helpers for the bank module.

// Accounts returns a reference to an account by ID, and whether it exists.
func Accounts(id int) (*Account, bool) {
    return GetAccount(id)
}

// History returns the current transaction log copy.
func History() []Transaction {
    return GetTransactions()
}



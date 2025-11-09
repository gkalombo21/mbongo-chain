package bank_test

import (
	"testing"

	"github.com/gkalombo21/mbongo-chain/internal/bank"
)

// TestCreateAccount_Success tests successful account creation with various inputs.
func TestCreateAccount_Success(t *testing.T) {
	tests := []struct {
		name           string
		owner          string
		initialBalance float64
		expectedOwner  string
		expectedBalance float64
	}{
		{
			name:            "Create account with positive balance",
			owner:           "Alice",
			initialBalance:  100.50,
			expectedOwner:   "Alice",
			expectedBalance: 100.50,
		},
		{
			name:            "Create account with zero balance",
			owner:           "Bob",
			initialBalance:  0.0,
			expectedOwner:   "Bob",
			expectedBalance: 0.0,
		},
		{
			name:            "Create account with negative balance (should be set to 0)",
			owner:           "Charlie",
			initialBalance:  -50.0,
			expectedOwner:   "Charlie",
			expectedBalance: 0.0,
		},
		{
			name:            "Create account with empty owner name (should default to unknown)",
			owner:           "",
			initialBalance:  200.0,
			expectedOwner:   "unknown",
			expectedBalance: 200.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := bank.CreateAccount(tt.owner, tt.initialBalance)
			
			if acct == nil {
				t.Fatal("CreateAccount returned nil")
			}
			
			if acct.OwnerName != tt.expectedOwner {
				t.Errorf("Expected owner %s, got %s", tt.expectedOwner, acct.OwnerName)
			}
			
			if acct.Balance != tt.expectedBalance {
				t.Errorf("Expected balance %.2f, got %.2f", tt.expectedBalance, acct.Balance)
			}
			
			if acct.ID <= 0 {
				t.Errorf("Expected positive account ID, got %d", acct.ID)
			}
			
			if acct.CreatedAt.IsZero() {
				t.Error("Expected CreatedAt to be set, got zero time")
			}
			
			t.Logf("Successfully created account #%d for %s with balance %.2f", acct.ID, acct.OwnerName, acct.Balance)
		})
	}
}

// TestCreateAccount_UniqueIDs tests that each account gets a unique ID.
func TestCreateAccount_UniqueIDs(t *testing.T) {
	acct1 := bank.CreateAccount("User1", 100.0)
	acct2 := bank.CreateAccount("User2", 200.0)
	acct3 := bank.CreateAccount("User3", 300.0)
	
	if acct1.ID == acct2.ID || acct1.ID == acct3.ID || acct2.ID == acct3.ID {
		t.Error("Expected all accounts to have unique IDs")
	}
	
	t.Logf("Created accounts with unique IDs: #%d, #%d, #%d", acct1.ID, acct2.ID, acct3.ID)
}

// TestDeposit_Success tests successful deposit operations.
func TestDeposit_Success(t *testing.T) {
	tests := []struct {
		name            string
		initialBalance  float64
		depositAmount   float64
		expectedBalance float64
	}{
		{
			name:            "Deposit positive amount to zero balance",
			initialBalance:  0.0,
			depositAmount:   50.0,
			expectedBalance: 50.0,
		},
		{
			name:            "Deposit positive amount to existing balance",
			initialBalance:  100.0,
			depositAmount:   25.75,
			expectedBalance: 125.75,
		},
		{
			name:            "Deposit large amount",
			initialBalance:  10.0,
			depositAmount:   1000.0,
			expectedBalance: 1010.0,
		},
		{
			name:            "Deposit small decimal amount",
			initialBalance:  100.0,
			depositAmount:   0.01,
			expectedBalance: 100.01,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := bank.CreateAccount("TestUser", tt.initialBalance)
			initialTxCount := len(bank.GetTransactions())
			
			acct.Deposit(tt.depositAmount)
			
			balance := acct.GetBalance()
			if balance != tt.expectedBalance {
				t.Errorf("Expected balance %.2f after deposit, got %.2f", tt.expectedBalance, balance)
			}
			
			// Verify transaction was recorded
			txCount := len(bank.GetTransactions())
			if txCount != initialTxCount+1 {
				t.Errorf("Expected transaction count to increase by 1, got %d transactions (was %d)", txCount, initialTxCount)
			}
			
			t.Logf("Successfully deposited %.2f, new balance: %.2f", tt.depositAmount, balance)
		})
	}
}

// TestDeposit_InvalidAmount tests that invalid deposit amounts are ignored.
func TestDeposit_InvalidAmount(t *testing.T) {
	tests := []struct {
		name           string
		initialBalance float64
		depositAmount  float64
		expectedBalance float64
	}{
		{
			name:            "Deposit zero amount (should be ignored)",
			initialBalance:  100.0,
			depositAmount:   0.0,
			expectedBalance: 100.0,
		},
		{
			name:            "Deposit negative amount (should be ignored)",
			initialBalance:  100.0,
			depositAmount:   -50.0,
			expectedBalance: 100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := bank.CreateAccount("TestUser", tt.initialBalance)
			initialTxCount := len(bank.GetTransactions())
			
			acct.Deposit(tt.depositAmount)
			
			balance := acct.GetBalance()
			if balance != tt.expectedBalance {
				t.Errorf("Expected balance to remain %.2f, got %.2f", tt.expectedBalance, balance)
			}
			
			// Verify no transaction was recorded for invalid deposits
			txCount := len(bank.GetTransactions())
			if txCount != initialTxCount {
				t.Errorf("Expected no new transaction for invalid deposit, but transaction count changed from %d to %d", initialTxCount, txCount)
			}
			
			t.Logf("Invalid deposit %.2f correctly ignored, balance remains %.2f", tt.depositAmount, balance)
		})
	}
}

// TestWithdraw_Success tests successful withdrawal operations.
func TestWithdraw_Success(t *testing.T) {
	tests := []struct {
		name            string
		initialBalance  float64
		withdrawAmount  float64
		expectedBalance float64
	}{
		{
			name:            "Withdraw exact balance",
			initialBalance:  100.0,
			withdrawAmount:  100.0,
			expectedBalance: 0.0,
		},
		{
			name:            "Withdraw partial balance",
			initialBalance:  100.0,
			withdrawAmount:  30.50,
			expectedBalance: 69.50,
		},
		{
			name:            "Withdraw small amount",
			initialBalance:  100.0,
			withdrawAmount:  0.01,
			expectedBalance: 99.99,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := bank.CreateAccount("TestUser", tt.initialBalance)
			initialTxCount := len(bank.GetTransactions())
			
			err := acct.Withdraw(tt.withdrawAmount)
			if err != nil {
				t.Errorf("Expected no error, got: %v", err)
			}
			
			balance := acct.GetBalance()
			if balance != tt.expectedBalance {
				t.Errorf("Expected balance %.2f after withdrawal, got %.2f", tt.expectedBalance, balance)
			}
			
			// Verify transaction was recorded
			txCount := len(bank.GetTransactions())
			if txCount != initialTxCount+1 {
				t.Errorf("Expected transaction count to increase by 1, got %d transactions (was %d)", txCount, initialTxCount)
			}
			
			t.Logf("Successfully withdrew %.2f, new balance: %.2f", tt.withdrawAmount, balance)
		})
	}
}

// TestWithdraw_InsufficientFunds tests withdrawal with insufficient funds.
func TestWithdraw_InsufficientFunds(t *testing.T) {
	acct := bank.CreateAccount("TestUser", 100.0)
	initialBalance := acct.GetBalance()
	initialTxCount := len(bank.GetTransactions())
	
	err := acct.Withdraw(150.0)
	if err == nil {
		t.Error("Expected error for insufficient funds, got nil")
	}
	
	balance := acct.GetBalance()
	if balance != initialBalance {
		t.Errorf("Expected balance to remain %.2f, got %.2f", initialBalance, balance)
	}
	
	// Verify no transaction was recorded for failed withdrawal
	txCount := len(bank.GetTransactions())
	if txCount != initialTxCount {
		t.Errorf("Expected no new transaction for failed withdrawal, but transaction count changed from %d to %d", initialTxCount, txCount)
	}
	
	t.Logf("Correctly rejected withdrawal of 150.0 from account with balance %.2f: %v", initialBalance, err)
}

// TestWithdraw_InvalidAmount tests withdrawal with invalid amounts.
func TestWithdraw_InvalidAmount(t *testing.T) {
	tests := []struct {
		name          string
		withdrawAmount float64
	}{
		{
			name:          "Withdraw zero amount",
			withdrawAmount: 0.0,
		},
		{
			name:          "Withdraw negative amount",
			withdrawAmount: -50.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := bank.CreateAccount("TestUser", 100.0)
			initialBalance := acct.GetBalance()
			initialTxCount := len(bank.GetTransactions())
			
			err := acct.Withdraw(tt.withdrawAmount)
			if err == nil {
				t.Error("Expected error for invalid withdrawal amount, got nil")
			}
			
			balance := acct.GetBalance()
			if balance != initialBalance {
				t.Errorf("Expected balance to remain %.2f, got %.2f", initialBalance, balance)
			}
			
			// Verify no transaction was recorded
			txCount := len(bank.GetTransactions())
			if txCount != initialTxCount {
				t.Errorf("Expected no new transaction for invalid withdrawal, but transaction count changed from %d to %d", initialTxCount, txCount)
			}
			
			t.Logf("Correctly rejected invalid withdrawal of %.2f: %v", tt.withdrawAmount, err)
		})
	}
}

// TestGetBalance tests balance retrieval.
func TestGetBalance(t *testing.T) {
	tests := []struct {
		name           string
		initialBalance float64
		deposits       []float64
		withdrawals    []float64
		expectedBalance float64
	}{
		{
			name:            "Get balance after creation",
			initialBalance:  100.0,
			deposits:        nil,
			withdrawals:     nil,
			expectedBalance: 100.0,
		},
		{
			name:            "Get balance after multiple deposits",
			initialBalance:  100.0,
			deposits:        []float64{50.0, 25.0, 10.0},
			withdrawals:     nil,
			expectedBalance: 185.0,
		},
		{
			name:            "Get balance after deposits and withdrawals",
			initialBalance:  100.0,
			deposits:        []float64{50.0, 25.0},
			withdrawals:     []float64{30.0, 20.0},
			expectedBalance: 125.0,
		},
		{
			name:            "Get balance after complex operations",
			initialBalance:  0.0,
			deposits:        []float64{100.0, 50.0, 25.0},
			withdrawals:     []float64{75.0},
			expectedBalance: 100.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			acct := bank.CreateAccount("TestUser", tt.initialBalance)
			
			for _, deposit := range tt.deposits {
				acct.Deposit(deposit)
			}
			
			for _, withdrawal := range tt.withdrawals {
				err := acct.Withdraw(withdrawal)
				if err != nil {
					t.Errorf("Unexpected error during withdrawal: %v", err)
				}
			}
			
			balance := acct.GetBalance()
			if balance != tt.expectedBalance {
				t.Errorf("Expected balance %.2f, got %.2f", tt.expectedBalance, balance)
			}
			
			t.Logf("Balance retrieved successfully: %.2f", balance)
		})
	}
}

// TestGetTransactions_History tests transaction history tracking.
func TestGetTransactions_History(t *testing.T) {
	// Clear any existing transactions by checking initial state
	initialTxCount := len(bank.GetTransactions())
	
	// Create account (should create 1 transaction)
	acct := bank.CreateAccount("TransactionTest", 100.0)
	txCount := len(bank.GetTransactions())
	if txCount != initialTxCount+1 {
		t.Errorf("Expected 1 transaction after account creation, got %d", txCount-initialTxCount)
	}
	t.Logf("Account creation transaction recorded: %d total transactions", txCount)
	
	// Deposit (should create 1 transaction)
	acct.Deposit(50.0)
	txCount = len(bank.GetTransactions())
	if txCount != initialTxCount+2 {
		t.Errorf("Expected 2 transactions after deposit, got %d", txCount-initialTxCount)
	}
	t.Logf("Deposit transaction recorded: %d total transactions", txCount)
	
	// Withdraw (should create 1 transaction)
	err := acct.Withdraw(25.0)
	if err != nil {
		t.Fatalf("Unexpected error during withdrawal: %v", err)
	}
	txCount = len(bank.GetTransactions())
	if txCount != initialTxCount+3 {
		t.Errorf("Expected 3 transactions after withdrawal, got %d", txCount-initialTxCount)
	}
	t.Logf("Withdrawal transaction recorded: %d total transactions", txCount)
	
	// Verify transaction types and amounts
	transactions := bank.GetTransactions()
	relevantTxs := transactions[initialTxCount:]
	
	if len(relevantTxs) < 3 {
		t.Fatalf("Expected at least 3 transactions, got %d", len(relevantTxs))
	}
	
	// Check create transaction
	if relevantTxs[0].Type != "create" {
		t.Errorf("Expected first transaction type 'create', got '%s'", relevantTxs[0].Type)
	}
	if relevantTxs[0].Amount != 100.0 {
		t.Errorf("Expected create transaction amount 100.0, got %.2f", relevantTxs[0].Amount)
	}
	
	// Check deposit transaction
	if relevantTxs[1].Type != "deposit" {
		t.Errorf("Expected second transaction type 'deposit', got '%s'", relevantTxs[1].Type)
	}
	if relevantTxs[1].Amount != 50.0 {
		t.Errorf("Expected deposit transaction amount 50.0, got %.2f", relevantTxs[1].Amount)
	}
	
	// Check withdraw transaction
	if relevantTxs[2].Type != "withdraw" {
		t.Errorf("Expected third transaction type 'withdraw', got '%s'", relevantTxs[2].Type)
	}
	if relevantTxs[2].Amount != 25.0 {
		t.Errorf("Expected withdraw transaction amount 25.0, got %.2f", relevantTxs[2].Amount)
	}
	
	t.Logf("All transaction types and amounts verified correctly")
}

// TestGetTransactions_HistoryFunction tests the History() wrapper function.
func TestGetTransactions_HistoryFunction(t *testing.T) {
	// Create account and perform operations
	acct := bank.CreateAccount("HistoryTest", 200.0)
	acct.Deposit(100.0)
	acct.Withdraw(50.0)
	
	// Get transactions via GetTransactions
	txs1 := bank.GetTransactions()
	
	// Get transactions via History wrapper
	txs2 := bank.History()
	
	// Both should return the same transactions
	if len(txs1) != len(txs2) {
		t.Errorf("Expected same transaction count, GetTransactions: %d, History: %d", len(txs1), len(txs2))
	}
	
	// Verify they contain the same data (checking last few transactions)
	if len(txs1) > 0 && len(txs2) > 0 {
		lastIdx := len(txs1) - 1
		if txs1[lastIdx].Type != txs2[lastIdx].Type || txs1[lastIdx].Amount != txs2[lastIdx].Amount {
			t.Error("GetTransactions and History returned different transaction data")
		}
	}
	
	t.Logf("History() function correctly returns same transactions as GetTransactions()")
}

// TestAccount_ConcurrentOperations tests basic concurrency safety.
func TestAccount_ConcurrentOperations(t *testing.T) {
	acct := bank.CreateAccount("ConcurrentTest", 1000.0)
	
	// Perform multiple concurrent deposits and withdrawals
	done := make(chan bool, 10)
	
	// Concurrent deposits
	for i := 0; i < 5; i++ {
		go func() {
			acct.Deposit(10.0)
			done <- true
		}()
	}
	
	// Concurrent withdrawals
	for i := 0; i < 5; i++ {
		go func() {
			acct.Withdraw(5.0)
			done <- true
		}()
	}
	
	// Wait for all operations to complete
	for i := 0; i < 10; i++ {
		<-done
	}
	
	// Expected: 1000 + (5 * 10) - (5 * 5) = 1000 + 50 - 25 = 1025
	expectedBalance := 1025.0
	balance := acct.GetBalance()
	
	if balance != expectedBalance {
		t.Errorf("Expected balance %.2f after concurrent operations, got %.2f", expectedBalance, balance)
	}
	
	t.Logf("Concurrent operations completed successfully, final balance: %.2f", balance)
}


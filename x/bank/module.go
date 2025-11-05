package bank

import (
	"fmt"

	"github.com/gkalombo21/mbongo-chain/x/bank/keeper"
)

// AppModule basic structure placeholder for compatibility
type AppModule struct{}

// Name returns the module name.
func (AppModule) Name() string {
	return "bank"
}

// RegisterTypes placeholder.
func (AppModule) RegisterTypes() {
	fmt.Println("Registering bank module types...")
}

// StartBankModule initializes the bank module
func StartBankModule() {
	keeper.StartBankModule()
}

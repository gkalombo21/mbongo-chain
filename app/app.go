package app

import (
	"fmt"

	"github.com/gkalombo21/mbongo-chain/x/bank"
)

// StartApp launches the core MBongo application
func StartApp() {
	fmt.Println("MBongo App core loaded.")

	// Initialize bank module
	bank.StartBankModule()
}

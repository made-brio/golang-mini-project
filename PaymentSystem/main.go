package main

import (
	"fmt"
	"payment-system/banking"
	"payment-system/logger"
)

func main() {
	account := &banking.Account{Balance: 1000.0}

	fmt.Printf("Initial Balance: %.2f\n", account.CheckBalance())

	// Test case 1: Successful transaction
	account.SafeTransfer(100.0, logger.Log)
	fmt.Printf("Balance after 100 transfer: %.2f\n", account.CheckBalance())

	// Test case 2: Insufficient funds
	account.SafeTransfer(2000.0, logger.Log)
	fmt.Printf("Balance after 2000 transfer attempt: %.2f\n", account.CheckBalance())

	// Test case 3: Negative transfer amount
	account.SafeTransfer(-50.0, logger.Log)
	fmt.Printf("Balance after -50 transfer attempt: %.2f\n", account.CheckBalance())
}

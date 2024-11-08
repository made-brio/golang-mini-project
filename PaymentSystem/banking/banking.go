package banking

import (
	"errors"
	"fmt"
	"payment-system/logger"
)

// Account represents a user account with balance
type Account struct {
	Balance float64
}

// CheckBalance checks the current balance
func (a *Account) CheckBalance() float64 {
	return a.Balance
}

// Transfer handles the transfer of funds between accounts
func (a *Account) Transfer(amount float64, logFunc func(string)) error {
	defer logger.Log("Transaction processed") // Ensure log entry is created

	// Handle negative transfer amount with panic
	if amount < 0 {
		panic("Transfer amount cannot be negative")
	}

	// Check if balance is sufficient
	if a.Balance < amount {
		return errors.New("insufficient balance")
	}

	// Deduct amount from balance
	a.Balance -= amount
	logFunc(fmt.Sprintf("Transferred: %0.2f", amount))
	return nil
}

// SafeTransfer uses recover to handle any panic from Transfer
func (a *Account) SafeTransfer(amount float64, logFunc func(string)) {
	defer func() {
		if r := recover(); r != nil {
			logFunc(fmt.Sprintf("Recovered from panic: %v", r))
		}
	}()

	err := a.Transfer(amount, logFunc)
	if err != nil {
		logFunc(fmt.Sprintf("Error: %v", err))
	}
}

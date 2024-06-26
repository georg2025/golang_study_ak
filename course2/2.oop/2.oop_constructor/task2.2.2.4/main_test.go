package main

import (
	"testing"
)

func TestNewCustomerWithNameAndAccount(t *testing.T) {
	savings := &SavingsAccount{}
	savings.Deposit(1000)
	customer := NewCustomer(1, WithName("John Doe"), WithAccount(savings))

	if customer.Name != "John Doe" {
		t.Errorf("expected customer name to be 'John Doe', got %v", customer.Name)
	}

	if customer.Account.Balance() != 1000 {
		t.Errorf("expected account balance to be 1000, got %v", customer.Account.Balance())
	}

	customer.Account.Withdraw(500)

	if customer.Account.Balance() != 500 {
		t.Errorf("expected account balance to be 500, got %v", customer.Account.Balance())
	}
}

func TestWithdrawFromCheckingAccountWithInsufficientBalance(t *testing.T) {
	checking := &CheckingAccount{}
	checking.Deposit(50)
	err := checking.Withdraw(100)

	if err == nil {
		t.Errorf("expected error when withdrawing more than balance, got nil")
	}

	if err.Error() != "not enogh money" {
		t.Errorf("expected error message 'not enogh money', got %v", err.Error())
	}

	if checking.Balance() != 50 {
		t.Errorf("expected balance to remain 50, got %v", checking.Balance())
	}
}

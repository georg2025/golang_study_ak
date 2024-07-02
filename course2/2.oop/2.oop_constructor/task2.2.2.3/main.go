package main

import (
	"fmt"
)

type Account interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CheckingAccount struct {
	balance float64
}

type SavingsAccount struct {
	CheckingAccount
	MinimumDeposit float64
}

func (c *CheckingAccount) Deposit(amount float64) error {
	c.balance += amount
	return nil
}

func (c *SavingsAccount) Deposit(amount float64) error {
	c.balance += amount
	return nil
}

func (c *CheckingAccount) Withdraw(amount float64) error {
	if c.balance < amount {
		return fmt.Errorf("not enogh money")
	}

	c.balance -= amount
	return nil
}

func (c *SavingsAccount) Withdraw(amount float64) error {
	if c.balance-amount < c.MinimumDeposit {
		return fmt.Errorf("your ballance cant be less than %v", c.MinimumDeposit)
	}

	c.balance -= amount
	return nil
}

func (c *CheckingAccount) Balance() float64 {
	return c.balance
}
func (c *SavingsAccount) Balance() float64 {
	return c.balance
}

type Customer struct {
	ID      int
	Name    string
	Account Account
}

type CustomerOption func(*Customer)

func WithName(name string) CustomerOption {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}

func NewCustomer(id int, options ...CustomerOption) *Customer {
	customer := &Customer{}

	for _, option := range options {
		option(customer)
	}

	return customer
}

func main() {
}

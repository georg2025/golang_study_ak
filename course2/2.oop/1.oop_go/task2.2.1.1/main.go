package main

import "fmt"

type Accounter interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
	Balance() float64
}

type CurrentAccount struct {
	balance float64
}

type SavingsAccount struct {
	CurrentAccount
	MinimumDeposit float64
}

func (c *CurrentAccount) Deposit(amount float64) error {
	c.balance += amount
	return nil
}

func (c *SavingsAccount) Deposit(amount float64) error {
	c.balance += amount
	return nil
}

func (c *CurrentAccount) Withdraw(amount float64) error {
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

func (c *CurrentAccount) Balance() float64 {
	return c.balance
}
func (c *SavingsAccount) Balance() float64 {
	return c.balance
}

func ProcessAccount(a Accounter) {
	a.Deposit(500)
	a.Withdraw(200)
	fmt.Printf("Balance: %.2f\n", a.Balance())
}

func main() {
	c := &CurrentAccount{}
	s := &SavingsAccount{MinimumDeposit: 500}
	ProcessAccount(c)
	ProcessAccount(s)
}

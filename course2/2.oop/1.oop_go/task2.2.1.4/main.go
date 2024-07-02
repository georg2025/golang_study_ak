package main

import (
	"fmt"
)

type PaymentMethod interface {
	Pay(amount float64) error
}

type CreditCard struct {
	balance float64
}

type Bitcoin struct {
	balance float64
}

func (b *Bitcoin) Pay(amount float64) error {
	if amount <= 0.0 {
		return fmt.Errorf("error with pay ammount")
	}

	if b.balance < amount {
		return fmt.Errorf("not enogh money")
	}

	b.balance -= amount
	return nil
}

func (card *CreditCard) Pay(amount float64) error {
	if amount <= 0.0 {
		return fmt.Errorf("error with pay ammount")
	}

	if card.balance < amount {
		return fmt.Errorf("not enogh money")
	}

	card.balance -= amount
	return nil
}

func ProcessPayment(p PaymentMethod, amount float64) {
	err := p.Pay(amount)

	if err != nil {
		fmt.Println("Cant proceed payment:", err)
	}
}

func main() {
	cc := &CreditCard{balance: 500.00}
	btc := &Bitcoin{balance: 2.00}
	ProcessPayment(cc, 200.00)
	ProcessPayment(btc, 1.00)
	fmt.Println(cc.balance)
	fmt.Println(btc.balance)
}

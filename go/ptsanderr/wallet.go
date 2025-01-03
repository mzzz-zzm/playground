package ptsanderr

import (
	"fmt"
	"errors"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

// this i/f is defined in fmt package
// type Stringer interface {
// 	String() string
// }

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrInsufficientFunds
	}

	w.balance = w.balance - amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
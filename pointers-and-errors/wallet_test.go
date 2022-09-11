package main

import (
	"testing"
)

func assertEquals(t testing.TB, expected any, received any) {
	t.Helper()
	if expected != received {
		t.Errorf("expected %v to equal %v", expected, received)
	}
}

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(0)}
		wallet.Deposit(Bitcoin(10))

		assertEquals(t, Bitcoin(10), wallet.Balance())
	})

	t.Run("withdraw more than balance", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(11))

		assertEquals(t, err, InsufficientFundsError)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertEquals(t, Bitcoin(5), wallet.Balance())
		assertEquals(t, err, nil)
	})
}

package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrAccountNotFound = errors.New("account not found")

	ErrAccountOriginNotFound = errors.New("account origin not found")

	ErrAccountDestinationNotFound = errors.New("account destination not found")

	ErrInsufficientBalance = errors.New("origin account does not have sufficient Balance")
)

type AccountID string

func (a AccountID) String() string {
	return string(a)
}

type Account struct {
	Id        AccountID `json:"id" example:"1"`
	Name      string    `json:"name" example:"Leo Messi"`
	Cpf       string    `json:"Cpf" example:"00.00.111.11"`
	Balance   Money     `json:"Balance" example:"40000"`
	CreatedAt time.Time `json:"created_at" example:"2019-11-09T21:21:46+00:00"`
}

type DeleteMessage struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type (
	AccountRepository interface {
		Create(context.Context, Account) (Account, error)
		UpdateBalance(context.Context, AccountID, Money) error
		FindAll(context.Context) ([]Account, error)
		FindByID(context.Context, AccountID) (Account, error)
		DeleteByID(context.Context, AccountID) (bool, error)
		FindBalance(context.Context, AccountID) (Account, error)
	}
)

func NewAccount(ID AccountID, name, CPF string, balance Money, createdAt time.Time) Account {
	return Account{
		Id:        ID,
		Name:      name,
		Cpf:       CPF,
		Balance:   balance,
		CreatedAt: createdAt,
	}
}

func (a *Account) Deposit(amount Money) {
	a.Balance += amount
}

func (a *Account) Withdraw(amount Money) error {
	if a.Balance < amount {
		return ErrInsufficientBalance
	}

	a.Balance -= amount

	return nil
}

func (a Account) ID() AccountID {
	return a.Id
}

func NewAccountBalance(balance Money) Account {
	return Account{Balance: balance}
}
func DeleteAccountBalance(success bool) bool {
	return success
}

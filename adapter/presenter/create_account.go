package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	"github.com/felixa1996/go-plate/usecase"
)

type createAccountPresenter struct{}

func NewCreateAccountPresenter() usecase.CreateAccountPresenter {
	return createAccountPresenter{}
}

func (a createAccountPresenter) Output(account domain.Account) usecase.CreateAccountOutput {
	return usecase.CreateAccountOutput{
		ID:        account.ID().String(),
		Name:      account.Name,
		CPF:       account.Cpf,
		Balance:   account.Balance.Float64(),
		CreatedAt: account.CreatedAt.Format(time.RFC3339),
	}
}

package presenter

import (
	"github.com/felixa1996/go-plate/usecase"
)

type deleteAccountBalancePresenter struct{}

func NewDeleteAccountBalancePresenter() usecase.DeleteAccountBalancePresenter {
	return deleteAccountBalancePresenter{}
}

func (a deleteAccountBalancePresenter) Output(success bool) usecase.DeleteAccountBalanceOutput {
	return usecase.DeleteAccountBalanceOutput{Success: success}
}

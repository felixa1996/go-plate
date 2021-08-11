package presenter

import (
	usecase "github.com/felixa1996/go-plate/usecase/receipt_lunar"
)

type deleteOneReceiptLunarPresenter struct{}

func NewDeleteOneReceiptLunarPresenter() usecase.DeleteOneReceiptLunarPresenter {
	return deleteOneReceiptLunarPresenter{}
}

func (a deleteOneReceiptLunarPresenter) Output(success bool, Id string) usecase.DeleteOneReceiptLunarOutput {
	return usecase.DeleteOneReceiptLunarOutput{Success: success, Id: Id}
}

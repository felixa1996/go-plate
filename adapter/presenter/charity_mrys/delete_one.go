package presenter

import (
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type deleteOneCharityMrysPresenter struct{}

func NewDeleteOneCharityMrysPresenter() usecase.DeleteOneCharityMrysPresenter {
	return deleteOneCharityMrysPresenter{}
}

func (a deleteOneCharityMrysPresenter) Output(success bool, Id string) usecase.DeleteOneCharityMrysOutput {
	return usecase.DeleteOneCharityMrysOutput{Success: success, Id: Id}
}

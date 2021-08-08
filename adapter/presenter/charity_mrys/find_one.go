package presenter

import (
	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type findCharityMrysPresenter struct{}

func NewFindCharityMrysPresenter() usecase.FindCharityMrysPresenter {
	return findCharityMrysPresenter{}
}

func (a findCharityMrysPresenter) Output(model domain.CharityMrys) usecase.FindCharityMrysOutput {
	return usecase.FindCharityMrysOutput{
		Id:          model.Id,
		Name:        model.Name,
		Amount:      model.Amount.Float64(),
		Month:       model.Month,
		Year:        model.Year,
		Description: model.Description,
		UserID:      model.UserID,
		Branch:      model.Branch,
		Username:    model.Username,
		CreatedAt:   model.CreatedAt,
	}
}

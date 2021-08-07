package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type createCharityMrysPresenter struct{}

func NewCreateCharityMrysPresenter() usecase.CreateCharityMrysPresenter {
	return createCharityMrysPresenter{}
}

func (a createCharityMrysPresenter) Output(model domain.CharityMrys) usecase.CreateCharityMrysOutput {
	return usecase.CreateCharityMrysOutput{
		Id:          model.Id,
		Name:        model.Name,
		Amount:      model.Amount,
		Year:        model.Year,
		Month:       model.Month,
		Description: model.Description,
		Branch:      model.Branch,
		CreatedAt:   model.CreatedAt.Format(time.RFC3339),
	}
}

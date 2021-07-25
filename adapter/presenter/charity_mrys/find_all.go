package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type findAllCharityMrysPresenter struct{}

func NewFindAllCharityMrysPresenter() usecase.FindAllCharityMrysPresenter {
	return findAllCharityMrysPresenter{}
}

func (a findAllCharityMrysPresenter) Output(result domain.CharityMrysAll) usecase.FindAllCharityMrysOutput {
	var o = usecase.FindAllCharityMrysOutput{}

	for _, model := range result.Data {
		o.Data = append(o.Data, usecase.FindAllCharityMrysData{
			ID:          model.Id,
			Name:        model.Name,
			Amount:      model.Amount.Float64(),
			Month:       model.Month,
			Year:        model.Year,
			Description: model.Description,
			CreatedAt:   model.CreatedAt.Format(time.RFC3339),
		})
	}

	return o
}

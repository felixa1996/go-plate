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

func (a findAllCharityMrysPresenter) Output(CharityMryss []domain.CharityMrys) []usecase.FindAllCharityMrysOutput {
	var o = make([]usecase.FindAllCharityMrysOutput, 0)

	for _, CharityMrys := range CharityMryss {
		o = append(o, usecase.FindAllCharityMrysOutput{
			ID:          CharityMrys.Id,
			Name:        CharityMrys.Name,
			Amount:      CharityMrys.Amount.Float64(),
			Month:       CharityMrys.Month,
			Year:        CharityMrys.Year,
			Description: CharityMrys.Description,
			CreatedAt:   CharityMrys.CreatedAt.Format(time.RFC3339),
		})
	}

	return o
}

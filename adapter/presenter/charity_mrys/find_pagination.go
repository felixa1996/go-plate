package presenter

import (
	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type findPaginationCharityMrysPresenter struct{}

func NewFindPaginationCharityMrysPresenter() usecase.FindPaginationCharityMrysPresenter {
	return findPaginationCharityMrysPresenter{}
}

func (a findPaginationCharityMrysPresenter) Output(result domain.CharityMrysPagination) usecase.FindPaginationCharityMrysOutput {
	var o = usecase.FindPaginationCharityMrysOutput{}

	for _, model := range result.Data {
		o.Data = append(o.Data, usecase.FindPaginationCharityMrysOutputData{
			Id:          model.Id,
			Name:        model.Name,
			Amount:      model.Amount.Float64(),
			Month:       model.Month,
			Year:        model.Year,
			Description: model.Description,
			CreatedAt:   model.CreatedAt,
		})
	}
	o.Meta = result.Meta

	return o
}

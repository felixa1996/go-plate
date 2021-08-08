package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type createBulkCharityMrysPresenter struct{}

func NewCreateBulkCharityMrysPresenter() usecase.CreateBulkCharityMrysPresenter {
	return createBulkCharityMrysPresenter{}
}

func (a createBulkCharityMrysPresenter) Output(model []domain.CharityMrys) []usecase.CreateBulkCharityMrysOutput {
	data := []usecase.CreateBulkCharityMrysOutput{}
	for _, detail := range model {
		d := usecase.CreateBulkCharityMrysOutput{
			Id:          detail.Id,
			Name:        detail.Name,
			Amount:      detail.Amount,
			Year:        detail.Year,
			Month:       detail.Month,
			Description: detail.Description,
			UserID:      detail.UserID,
			Username:    detail.Username,
			Branch:      detail.Branch,
			CreatedAt:   detail.CreatedAt.Format(time.RFC3339),
		}
		data = append(data, d)
	}
	return data
}

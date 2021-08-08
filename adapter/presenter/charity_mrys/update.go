package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type updateCharityMrysPresenter struct{}

func NewUpdateCharityMrysPresenter() usecase.UpdateCharityMrysPresenter {
	return updateCharityMrysPresenter{}
}

func (a updateCharityMrysPresenter) Output(model domain.CharityMrys) usecase.UpdateCharityMrysOutput {
	return usecase.UpdateCharityMrysOutput{
		Id:          model.Id,
		Name:        model.Name,
		Amount:      model.Amount,
		Year:        model.Year,
		Month:       model.Month,
		Description: model.Description,
		Branch:      model.Branch,
		UserID:      model.UserID,
		Username:    model.Username,
		CreatedAt:   model.CreatedAt.Format(time.RFC3339),
	}
}

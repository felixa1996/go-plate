package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/receipt_lunar"
)

type createReceiptLunarPresenter struct{}

func NewCreateReceiptLunarPresenter() usecase.CreateReceiptLunarPresenter {
	return createReceiptLunarPresenter{}
}

func (a createReceiptLunarPresenter) Output(model domain.ReceiptLunar) usecase.CreateReceiptLunarOutput {
	return usecase.CreateReceiptLunarOutput{
		Id:                 model.Id,
		InternationalDate:  model.InternationalDate,
		LunarDate:          model.LunarDate,
		Description:        model.Description,
		Total:              model.Total,
		UserID:             model.UserID,
		Username:           model.Username,
		Branch:             model.Branch,
		ReceiptLunarDetail: model.ReceiptLunarDetail,
		CreatedAt:          model.CreatedAt.Format(time.RFC3339),
	}
}

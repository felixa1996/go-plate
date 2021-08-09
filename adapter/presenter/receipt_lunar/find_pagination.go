package presenter

import (
	"time"

	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/receipt_lunar"
)

type findPaginationReceiptLunarPresenter struct{}

func NewFindPaginationReceiptLunarPresenter() usecase.FindPaginationReceiptLunarPresenter {
	return findPaginationReceiptLunarPresenter{}
}

func (a findPaginationReceiptLunarPresenter) Output(result domain.ReceiptLunarPagination) usecase.FindPaginationReceiptLunarOutput {
	var o = usecase.FindPaginationReceiptLunarOutput{}

	for _, model := range result.Data {
		o.Data = append(o.Data, usecase.FindPaginationReceiptLunarOutputData{
			Id:                model.Id,
			InternationalDate: model.InternationalDate,
			LunarDate:         model.LunarDate,
			Description:       model.Description,
			Total:             model.Total,
			UserID:            model.UserID,
			Username:          model.Username,
			Branch:            model.Branch,
			CreatedAt:         model.CreatedAt.Format(time.RFC3339),
		})
	}
	o.Meta = result.Meta

	return o
}

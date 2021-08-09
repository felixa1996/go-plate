package action

import (
	"net/http"
	"strconv"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	usecase "github.com/felixa1996/go-plate/usecase/receipt_lunar"
)

type FindPaginationReceiptLunarAction struct {
	uc  usecase.FindPaginationReceiptLunarUseCase
	log logger.Logger
}

func NewFindPaginationReceiptLunarAction(uc usecase.FindPaginationReceiptLunarUseCase, log logger.Logger) FindPaginationReceiptLunarAction {
	return FindPaginationReceiptLunarAction{
		uc:  uc,
		log: log,
	}
}

func (a FindPaginationReceiptLunarAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_pagination_receipt_lunar"

	currentPage, err := strconv.Atoi(r.URL.Query().Get("currentPage"))
	if err != nil {
		logging.NewError(a.log, err, logKey, http.StatusBadRequest).Log("error when parse currentPage")
	}
	perPage, err := strconv.Atoi(r.URL.Query().Get("perPage"))
	if err != nil {
		logging.NewError(a.log, err, logKey, http.StatusBadRequest).Log("error when parse perPage")
	}
	var sort = r.URL.Query().Get("sort")
	var search = r.URL.Query().Get("search")

	output, err := a.uc.Execute(r.Context(), currentPage, perPage, sort, search)
	if err != nil {
		logging.NewError(a.log, err, logKey, http.StatusInternalServerError).Log("error when returning ReceiptLunar list")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning ReceiptLunar list")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

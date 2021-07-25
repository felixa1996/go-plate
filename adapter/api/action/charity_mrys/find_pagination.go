package action

import (
	"net/http"
	"strconv"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type FindPaginationCharityMrysAction struct {
	uc  usecase.FindPaginationCharityMrysUseCase
	log logger.Logger
}

func NewFindPaginationCharityMrysAction(uc usecase.FindPaginationCharityMrysUseCase, log logger.Logger) FindPaginationCharityMrysAction {
	return FindPaginationCharityMrysAction{
		uc:  uc,
		log: log,
	}
}

func (a FindPaginationCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_pagination_CharityMrys"

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
		logging.NewError(a.log, err, logKey, http.StatusInternalServerError).Log("error when returning CharityMrys list")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning CharityMrys list")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

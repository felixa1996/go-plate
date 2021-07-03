package action

import (
	"net/http"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type FindCharityMrysAction struct {
	uc  usecase.FindCharityMrysUseCase
	log logger.Logger
}

func NewFindCharityMrysAction(uc usecase.FindCharityMrysUseCase, log logger.Logger) FindCharityMrysAction {
	return FindCharityMrysAction{
		uc:  uc,
		log: log,
	}
}

func (a FindCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_one_charity_mrys"

	var id = r.URL.Query().Get("id")

	output, err := a.uc.Execute(r.Context(), domain.CharityMrysID(id))
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusBadRequest,
			).Log("error fetching one charity mrys")

			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		default:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error when returning one charity mrys")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning one charity mrys")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

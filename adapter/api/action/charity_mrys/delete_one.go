package action

import (
	"fmt"
	"net/http"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type DeleteOneCharityMrysAction struct {
	uc  usecase.DeleteOneCharityMrysUseCase
	log logger.Logger
}

func NewDeleteOneCharityMrysAction(uc usecase.DeleteOneCharityMrysUseCase, log logger.Logger) DeleteOneCharityMrysAction {
	return DeleteOneCharityMrysAction{
		uc:  uc,
		log: log,
	}
}

func (a DeleteOneCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "delete_one_charity_mrys"

	var id = r.URL.Query().Get("id")
	fmt.Print("Text " + id)

	output, err := a.uc.Execute(r.Context(), id)
	if err != nil {
		switch err {
		case domain.ErrCharityMrysNotFound:
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
			).Log("error when returning charity mrys")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning charity mrys")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

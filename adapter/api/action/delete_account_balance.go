package action

import (
	"fmt"
	"net/http"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/domain"
	"github.com/felixa1996/go-plate/usecase"
)

type DeleteAccountBalanceAction struct {
	uc  usecase.DeleteAccountBalanceUseCase
	log logger.Logger
}

func NewDeleteAccountBalanceAction(uc usecase.DeleteAccountBalanceUseCase, log logger.Logger) DeleteAccountBalanceAction {
	return DeleteAccountBalanceAction{
		uc:  uc,
		log: log,
	}
}

func (a DeleteAccountBalanceAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "delete_balance_account"

	var id = r.URL.Query().Get("id")
	fmt.Print("Text " + id)
	if !domain.IsValidUUID(id) {
		var err = response.ErrParameterInvalid
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("invalid parameter")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), domain.AccountID(id))
	if err != nil {
		switch err {
		case domain.ErrAccountNotFound:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusBadRequest,
			).Log("error fetching account balance")

			response.NewError(err, http.StatusBadRequest).Send(w)
			return
		default:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusInternalServerError,
			).Log("error when returning account balance")

			response.NewError(err, http.StatusInternalServerError).Send(w)
			return
		}
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning account balance")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

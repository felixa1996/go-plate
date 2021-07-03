package action

import (
	"net/http"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type FindAllCharityMrysAction struct {
	uc  usecase.FindAllCharityMrysUseCase
	log logger.Logger
}

func NewFindAllCharityMrysAction(uc usecase.FindAllCharityMrysUseCase, log logger.Logger) FindAllCharityMrysAction {
	return FindAllCharityMrysAction{
		uc:  uc,
		log: log,
	}
}

func (a FindAllCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "find_all_CharityMrys"

	output, err := a.uc.Execute(r.Context())
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when returning CharityMrys list")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning CharityMrys list")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

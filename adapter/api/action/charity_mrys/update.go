package action

import (
	"encoding/json"
	"net/http"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/adapter/validator"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type UpdateCharityMrysAction struct {
	uc        usecase.UpdateCharityMrysUseCase
	log       logger.Logger
	validator validator.Validator
}

func NewUpdateCharityMrysAction(uc usecase.UpdateCharityMrysUseCase, log logger.Logger, v validator.Validator) UpdateCharityMrysAction {
	return UpdateCharityMrysAction{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

func (a UpdateCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "update_charity_mrys"

	var id = r.URL.Query().Get("id")

	var input usecase.UpdateCharityMrysInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusBadRequest,
		).Log("error when decoding json")

		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	if errs := a.validateInput(input); len(errs) > 0 {
		logging.NewError(
			a.log,
			response.ErrInvalidInput,
			logKey,
			http.StatusBadRequest,
		).Log("invalid input")

		response.NewErrorMessage(errs, http.StatusBadRequest).Send(w)
		return
	}

	output, err := a.uc.Execute(r.Context(), input, id)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when updating a new charity_mrys")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success updating charity_mrys")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func (a UpdateCharityMrysAction) validateInput(input usecase.UpdateCharityMrysInput) []string {
	var msgs []string

	err := a.validator.Validate(input)
	if err != nil {
		for _, msg := range a.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}

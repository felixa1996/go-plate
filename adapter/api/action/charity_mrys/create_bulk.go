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

type CreateBulkCharityMrysAction struct {
	uc        usecase.CreateBulkCharityMrysUseCase
	log       logger.Logger
	validator validator.Validator
}

func NewCreateBulkCharityMrysAction(uc usecase.CreateBulkCharityMrysUseCase, log logger.Logger, v validator.Validator) CreateBulkCharityMrysAction {
	return CreateBulkCharityMrysAction{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

func (a CreateBulkCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_bulk_charity_mrys"

	var input usecase.CreateBulkCharityMrysInput
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

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		logging.NewError(
			a.log,
			err,
			logKey,
			http.StatusInternalServerError,
		).Log("error when creating a new charity_mrys")

		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}
	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success creating bulk charity_mrys")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func (a CreateBulkCharityMrysAction) validateInput(input usecase.CreateBulkCharityMrysInput) []string {
	var msgs []string

	err := a.validator.Validate(input)
	if err != nil {
		for _, msg := range a.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}

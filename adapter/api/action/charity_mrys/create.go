package action

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	gouuid "github.com/satori/go.uuid"

	"github.com/felixa1996/go-plate/adapter/api/logging"
	"github.com/felixa1996/go-plate/adapter/api/response"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/adapter/validator"
	"github.com/felixa1996/go-plate/infrastructure/broker"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

type CreateCharityMrysAction struct {
	uc        usecase.CreateCharityMrysUseCase
	log       logger.Logger
	validator validator.Validator
}

func NewCreateCharityMrysAction(uc usecase.CreateCharityMrysUseCase, log logger.Logger, v validator.Validator) CreateCharityMrysAction {
	return CreateCharityMrysAction{
		uc:        uc,
		log:       log,
		validator: v,
	}
}

func (a CreateCharityMrysAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "create_charity_mrys"

	var input usecase.CreateCharityMrysInput
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

	a.KafkaSendProducer(output)

	logging.NewInfo(a.log, logKey, http.StatusCreated).Log("success creating charity_mrys")

	response.NewSuccess(output, http.StatusCreated).Send(w)
}

func (a CreateCharityMrysAction) validateInput(input usecase.CreateCharityMrysInput) []string {
	var msgs []string

	err := a.validator.Validate(input)
	if err != nil {
		for _, msg := range a.validator.Messages() {
			msgs = append(msgs, msg)
		}
	}

	return msgs
}

func (a CreateCharityMrysAction) KafkaSendProducer(result ...interface{}) {

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := &broker.KafkaProducer{
		Ctx:    context.Background(),
		Log:    a.log,
		LogKey: logKey,
		Topic:  "charity_mrys_insert",
		Key:    gouuid.NewV4().String(),
		Value:  string(b),
	}

	go broker.Produce(t)
}

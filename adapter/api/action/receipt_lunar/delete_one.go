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
	"github.com/felixa1996/go-plate/domain"
	"github.com/felixa1996/go-plate/infrastructure/broker"
	usecase "github.com/felixa1996/go-plate/usecase/receipt_lunar"
)

type DeleteOneReceiptLunarAction struct {
	uc  usecase.DeleteOneReceiptLunarUseCase
	log logger.Logger
}

func NewDeleteOneReceiptLunarAction(uc usecase.DeleteOneReceiptLunarUseCase, log logger.Logger) DeleteOneReceiptLunarAction {
	return DeleteOneReceiptLunarAction{
		uc:  uc,
		log: log,
	}
}

func (a DeleteOneReceiptLunarAction) Execute(w http.ResponseWriter, r *http.Request) {
	const logKey = "delete_one_receipt_lunar"

	var id = r.URL.Query().Get("id")
	fmt.Print("Text " + id)

	output, err := a.uc.Execute(r.Context(), id)
	if err != nil {
		switch err {
		case domain.ErrReceiptLunarNotFound:
			logging.NewError(
				a.log,
				err,
				logKey,
				http.StatusBadRequest,
			).Log("error fetching one receipt_lunar")

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

	a.KafkaSendProducer(output)

	logging.NewInfo(a.log, logKey, http.StatusOK).Log("success when returning charity mrys")

	response.NewSuccess(output, http.StatusOK).Send(w)
}

func (a DeleteOneReceiptLunarAction) KafkaSendProducer(result ...interface{}) {

	b, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}

	t := &broker.KafkaProducer{
		Ctx:    context.Background(),
		Log:    a.log,
		LogKey: logKey,
		Topic:  "receipt_lunar_delete",
		Key:    gouuid.NewV4().String(),
		Value:  string(b),
	}

	go broker.Produce(t)
}

package route

import (
	"time"

	action "github.com/felixa1996/go-plate/adapter/api/action/receipt_lunar"
	"github.com/felixa1996/go-plate/adapter/logger"
	presenter "github.com/felixa1996/go-plate/adapter/presenter/receipt_lunar"
	"github.com/felixa1996/go-plate/adapter/repository"
	"github.com/felixa1996/go-plate/adapter/validator"
	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/receipt_lunar"
)

func ReceiptLunarFindPagination(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, auth *domain.UserJwt) action.FindPaginationReceiptLunarAction {
	var (
		uc = usecase.NewFindPaginationReceiptLunarInteractor(
			repository.NewReceiptLunarSQL(db),
			presenter.NewFindPaginationReceiptLunarPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewFindPaginationReceiptLunarAction(uc, log)
	)
	return act
}

func ReceiptLunarFindOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, auth *domain.UserJwt) action.FindReceiptLunarAction {
	var (
		uc = usecase.NewFindReceiptLunarInteractor(
			repository.NewReceiptLunarSQL(db),
			presenter.NewFindReceiptLunarPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewFindReceiptLunarAction(uc, log)
	)
	return act
}

// func ReceiptLunarDeleteOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration) action.DeleteOneReceiptLunarAction {
// 	var (
// 		uc = usecase.NewDeleteOneReceiptLunarInteractor(
// 			repository.NewReceiptLunarSQL(db),
// 			presenter.NewDeleteOneReceiptLunarPresenter(),
// 			ctxTimeout,
// 		)
// 		act = action.NewDeleteOneReceiptLunarAction(uc, log)
// 	)
// 	return act
// }

func ReceiptLunarCreateOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator, auth *domain.UserJwt) action.CreateReceiptLunarAction {
	var (
		uc = usecase.NewCreateReceiptLunarInteractor(
			repository.NewReceiptLunarSQL(db),
			presenter.NewCreateReceiptLunarPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewCreateReceiptLunarAction(uc, log, validator)
	)
	return act
}

// func ReceiptLunarUpdateOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator, auth *domain.UserJwt) action.UpdateReceiptLunarAction {
// 	var (
// 		uc = usecase.NewUpdateReceiptLunarInteractor(
// 			repository.NewReceiptLunarSQL(db),
// 			presenter.NewUpdateReceiptLunarPresenter(),
// 			auth,
// 			ctxTimeout,
// 		)
// 		act = action.NewUpdateReceiptLunarAction(uc, log, validator)
// 	)
// 	return act
// }

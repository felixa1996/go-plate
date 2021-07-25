package route

import (
	"time"

	action "github.com/felixa1996/go-plate/adapter/api/action/charity_mrys"
	"github.com/felixa1996/go-plate/adapter/logger"
	presenter "github.com/felixa1996/go-plate/adapter/presenter/charity_mrys"
	"github.com/felixa1996/go-plate/adapter/repository"
	"github.com/felixa1996/go-plate/adapter/validator"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

func CharityMrysFindAll(db repository.SQL, log logger.Logger, ctxTimeout time.Duration) action.FindAllCharityMrysAction {
	var (
		uc = usecase.NewFindAllCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewFindAllCharityMrysPresenter(),
			ctxTimeout,
		)
		act = action.NewFindAllCharityMrysAction(uc, log)
	)
	return act
}

func CharityMrysFindOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration) action.FindCharityMrysAction {
	var (
		uc = usecase.NewFindCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewFindCharityMrysPresenter(),
			ctxTimeout,
		)
		act = action.NewFindCharityMrysAction(uc, log)
	)
	return act
}

func CharityMrysDeleteOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration) action.DeleteOneCharityMrysAction {
	var (
		uc = usecase.NewDeleteOneCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewDeleteOneCharityMrysPresenter(),
			ctxTimeout,
		)
		act = action.NewDeleteOneCharityMrysAction(uc, log)
	)
	return act
}

func CharityMrysCreateOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator) action.CreateCharityMrysAction {
	var (
		uc = usecase.NewCreateCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewCreateCharityMrysPresenter(),
			ctxTimeout,
		)
		act = action.NewCreateCharityMrysAction(uc, log, validator)
	)
	return act
}

func CharityMrysCreateBulk(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator) action.CreateBulkCharityMrysAction {
	var (
		uc = usecase.NewCreateBulkCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewCreateBulkCharityMrysPresenter(),
			ctxTimeout,
		)
		act = action.NewCreateBulkCharityMrysAction(uc, log, validator)
	)
	return act
}

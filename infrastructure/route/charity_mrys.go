package route

import (
	"time"

	action "github.com/felixa1996/go-plate/adapter/api/action/charity_mrys"
	"github.com/felixa1996/go-plate/adapter/logger"
	presenter "github.com/felixa1996/go-plate/adapter/presenter/charity_mrys"
	"github.com/felixa1996/go-plate/adapter/repository"
	"github.com/felixa1996/go-plate/adapter/validator"
	"github.com/felixa1996/go-plate/domain"
	usecase "github.com/felixa1996/go-plate/usecase/charity_mrys"
)

func CharityMrysFindAll(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, auth *domain.UserJwt) action.FindAllCharityMrysAction {
	var (
		uc = usecase.NewFindAllCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewFindAllCharityMrysPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewFindAllCharityMrysAction(uc, log)
	)
	return act
}

func CharityMrysFindPagination(db repository.SQL, log logger.Logger, ctxTimeout time.Duration) action.FindPaginationCharityMrysAction {
	var (
		uc = usecase.NewFindPaginationCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewFindPaginationCharityMrysPresenter(),
			ctxTimeout,
		)
		act = action.NewFindPaginationCharityMrysAction(uc, log)
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

func CharityMrysCreateOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator, auth *domain.UserJwt) action.CreateCharityMrysAction {
	var (
		uc = usecase.NewCreateCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewCreateCharityMrysPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewCreateCharityMrysAction(uc, log, validator)
	)
	return act
}

func CharityMrysUpdateOne(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator, auth *domain.UserJwt) action.UpdateCharityMrysAction {
	var (
		uc = usecase.NewUpdateCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewUpdateCharityMrysPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewUpdateCharityMrysAction(uc, log, validator)
	)
	return act
}

func CharityMrysCreateBulk(db repository.SQL, log logger.Logger, ctxTimeout time.Duration, validator validator.Validator, auth *domain.UserJwt) action.CreateBulkCharityMrysAction {
	var (
		uc = usecase.NewCreateBulkCharityMrysInteractor(
			repository.NewCharityMrysSQL(db),
			presenter.NewCreateBulkCharityMrysPresenter(),
			auth,
			ctxTimeout,
		)
		act = action.NewCreateBulkCharityMrysAction(uc, log, validator)
	)
	return act
}

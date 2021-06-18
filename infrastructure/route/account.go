package route

import (
	"github.com/felixa1996/go-plate/adapter/api/action"
	"github.com/felixa1996/go-plate/adapter/logger"
	"github.com/felixa1996/go-plate/adapter/presenter"
	"github.com/felixa1996/go-plate/adapter/repository"
	"github.com/felixa1996/go-plate/usecase"
	"time"
)

func AccountFindAll(db repository.SQL, log logger.Logger, ctxTimeout time.Duration) action.FindAllAccountAction {
	var (
		uc = usecase.NewFindAllAccountInteractor(
			repository.NewAccountSQL(db),
			presenter.NewFindAllAccountPresenter(),
			ctxTimeout,
		)
		act = action.NewFindAllAccountAction(uc, log)
	)
	return act
}

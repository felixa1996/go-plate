package route

import (
	"time"

	action "github.com/felixa1996/go-plate/adapter/api/action/charity_mrys"
	"github.com/felixa1996/go-plate/adapter/logger"
	presenter "github.com/felixa1996/go-plate/adapter/presenter/charity_mrys"
	"github.com/felixa1996/go-plate/adapter/repository"
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

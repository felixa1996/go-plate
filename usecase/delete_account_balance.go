package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// DeleteAccountBalanceUseCase input port
	DeleteAccountBalanceUseCase interface {
		Execute(context.Context, domain.AccountID) (DeleteAccountBalanceOutput, error)
	}

	// DeleteAccountBalaceInput input data
	DeleteAccountBalaceInput struct {
		ID int64 `json:"balance" validate:"gt=0,required"`
	}

	// DeleteAccountBalancePresenter output port
	DeleteAccountBalancePresenter interface {
		Output(bool) DeleteAccountBalanceOutput
	}

	// FindAccountBalanceOutput output data
	DeleteAccountBalanceOutput struct {
		Success bool `json:"success"`
	}

	deleteBalanceAccountInteractor struct {
		repo       domain.AccountRepository
		presenter  DeleteAccountBalancePresenter
		ctxTimeout time.Duration
	}
)

// NewDeleteBalanceAccountInteractor creates new deleteBalanceAccountInteractor with its dependencies
func NewDeleteBalanceAccountInteractor(
	repo domain.AccountRepository,
	presenter DeleteAccountBalancePresenter,
	t time.Duration,
) DeleteAccountBalanceUseCase {
	return deleteBalanceAccountInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a deleteBalanceAccountInteractor) Execute(ctx context.Context, ID domain.AccountID) (DeleteAccountBalanceOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	account, err := a.repo.DeleteByID(ctx, ID)
	if err != nil {
		return a.presenter.Output(false), err
	}

	return a.presenter.Output(account), nil
}

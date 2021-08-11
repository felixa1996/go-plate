package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	DeleteOneReceiptLunarUseCase interface {
		Execute(context.Context, string) (DeleteOneReceiptLunarOutput, error)
	}

	DeleteOneReceiptLunarInput struct {
		ID string `json:"id" validate:"gt=0,required"`
	}

	DeleteOneReceiptLunarPresenter interface {
		Output(bool, string) DeleteOneReceiptLunarOutput
	}

	DeleteOneReceiptLunarOutput struct {
		Success bool   `json:"success"`
		Id      string `json:"id"`
	}

	deleteOneReceiptLunarInteractor struct {
		repo       domain.ReceiptLunarRepository
		presenter  DeleteOneReceiptLunarPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

func NewDeleteOneReceiptLunarInteractor(
	repo domain.ReceiptLunarRepository,
	presenter DeleteOneReceiptLunarPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) DeleteOneReceiptLunarUseCase {
	return deleteOneReceiptLunarInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a deleteOneReceiptLunarInteractor) Execute(ctx context.Context, ID string) (DeleteOneReceiptLunarOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	model, err := a.repo.DeleteByID(ctx, ID)
	if err != nil {
		return a.presenter.Output(false, ID), err
	}

	return a.presenter.Output(model, ID), nil
}

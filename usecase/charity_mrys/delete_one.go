package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// DeleteOneCharityMrysUseCase input port
	DeleteOneCharityMrysUseCase interface {
		Execute(context.Context, string) (DeleteOneCharityMrysOutput, error)
	}

	// DeleteOneCharityMrysInput input data
	DeleteOneCharityMrysInput struct {
		ID string `json:"id" validate:"gt=0,required"`
	}

	// DeleteOneCharityMrysPresenter output port
	DeleteOneCharityMrysPresenter interface {
		Output(bool) DeleteOneCharityMrysOutput
	}

	// FindOneCharityMrysOutput output data
	DeleteOneCharityMrysOutput struct {
		Success bool `json:"success"`
	}

	deleteOneCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  DeleteOneCharityMrysPresenter
		ctxTimeout time.Duration
	}
)

func NewDeleteOneCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter DeleteOneCharityMrysPresenter,
	t time.Duration,
) DeleteOneCharityMrysUseCase {
	return deleteOneCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a deleteOneCharityMrysInteractor) Execute(ctx context.Context, ID string) (DeleteOneCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	model, err := a.repo.DeleteByID(ctx, ID)
	if err != nil {
		return a.presenter.Output(false), err
	}

	return a.presenter.Output(model), nil
}

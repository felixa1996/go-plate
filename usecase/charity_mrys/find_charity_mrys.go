package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// FindCharityMrysUseCase input port
	FindCharityMrysUseCase interface {
		Execute(context.Context, string) (FindCharityMrysOutput, error)
	}

	// FindCharityMrysInput input data
	FindCharityMrysInput struct {
		ID int64 `json:"id" validate:"gt=0,required"`
	}

	// FindCharityMrysPresenter output port
	FindCharityMrysPresenter interface {
		Output(domain.CharityMrys) FindCharityMrysOutput
	}

	// FindCharityMrysOutput output data
	FindCharityMrysOutput struct {
		Id          string    `json:"id"`
		Name        string    `json:"name"`
		Amount      float64   `json:"Amount"`
		Month       int32     `json:"Month"`
		Year        int32     `json:"Year"`
		Description string    `json:"Description"`
		CreatedAt   time.Time `json:"CreatedAt"`
	}

	findCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  FindCharityMrysPresenter
		ctxTimeout time.Duration
	}
)

// NewFindCharityMrysInteractor creates new findCharityMrysInteractor with its dependencies
func NewFindCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter FindCharityMrysPresenter,
	t time.Duration,
) FindCharityMrysUseCase {
	return findCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a findCharityMrysInteractor) Execute(ctx context.Context, ID string) (FindCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	model, err := a.repo.FindByID(ctx, ID)
	if err != nil {
		return a.presenter.Output(domain.CharityMrys{}), err
	}

	return a.presenter.Output(model), nil
}

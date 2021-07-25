package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// FindAllCharityMrysUseCase input port
	FindAllCharityMrysUseCase interface {
		Execute(context.Context) ([]FindAllCharityMrysOutput, error)
	}

	// FindAllCharityMrysPresenter output port
	FindAllCharityMrysPresenter interface {
		Output([]domain.CharityMrys) []FindAllCharityMrysOutput
	}

	// FindAllCharityMrysOutput outputData
	FindAllCharityMrysOutput struct {
		ID          string  `json:"id"`
		Name        string  `json:"name"`
		Amount      float64 `json:"amount"`
		Month       int32   `json:"month"`
		Year        int32   `json:"year"`
		Description string  `json:"description"`
		CreatedAt   string  `json:"created_at"`
	}

	findAllCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  FindAllCharityMrysPresenter
		ctxTimeout time.Duration
	}
)

// NewFindAllCharityMrysInteractor creates new findAllCharityMrysInteractor with its dependencies
func NewFindAllCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter FindAllCharityMrysPresenter,
	t time.Duration,
) FindAllCharityMrysUseCase {
	return findAllCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a findAllCharityMrysInteractor) Execute(ctx context.Context) ([]FindAllCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	CharityMryss, err := a.repo.FindAll(ctx)
	if err != nil {
		return a.presenter.Output([]domain.CharityMrys{}), err
	}

	return a.presenter.Output(CharityMryss), nil
}

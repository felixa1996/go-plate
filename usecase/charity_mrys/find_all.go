package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// FindAllCharityMrysUseCase input port
	FindAllCharityMrysUseCase interface {
		Execute(context.Context) (FindAllCharityMrysOutput, error)
	}

	// FindAllCharityMrysPresenter output port
	FindAllCharityMrysPresenter interface {
		Output(domain.CharityMrysAll) FindAllCharityMrysOutput
	}

	FindAllCharityMrysOutput struct {
		Data []FindAllCharityMrysData `json:"data"`
	}

	findAllCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  FindAllCharityMrysPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

// NewFindAllCharityMrysInteractor creates new findAllCharityMrysInteractor with its dependencies
func NewFindAllCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter FindAllCharityMrysPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) FindAllCharityMrysUseCase {
	return findAllCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a findAllCharityMrysInteractor) Execute(ctx context.Context) (FindAllCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	result, err := a.repo.FindAll(ctx, a.auth)
	if err != nil {
		return a.presenter.Output(domain.CharityMrysAll{}), err
	}

	return a.presenter.Output(result), nil
}

type FindAllCharityMrysData struct {
	ID          string        `json:"id"`
	Name        string        `json:"name"`
	Amount      float64       `json:"amount"`
	Month       int32         `json:"month"`
	Year        int32         `json:"year"`
	Description string        `json:"description"`
	UserID      string        `json:"user_id"`
	Username    string        `json:"username"`
	Branch      domain.Branch `json:"branch"`
	CreatedAt   string        `json:"created_at"`
}

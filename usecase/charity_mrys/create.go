package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// CreateCharityMrysUseCase input port
	CreateCharityMrysUseCase interface {
		Execute(context.Context, CreateCharityMrysInput) (CreateCharityMrysOutput, error)
	}

	// CreateCharityMrysInput input data
	CreateCharityMrysInput struct {
		Id          string        `json:"id"`
		Name        string        `json:"name" validate:"required"`
		Amount      int32         `json:"amount" validate:"required"`
		Month       int32         `json:"month" validate:"required"`
		Year        int32         `json:"year" validate:"required"`
		Description string        `json:"description"`
		Branch      domain.Branch `json:"branch"`
	}

	// CreateCharityMrysPresenter output port
	CreateCharityMrysPresenter interface {
		Output(domain.CharityMrys) CreateCharityMrysOutput
	}

	// CreateCharityMrysOutput output data
	CreateCharityMrysOutput struct {
		Id          string        `json:"id"`
		Name        string        `json:"name" validate:"required"`
		Amount      domain.Money  `json:"amount" validate:"required"`
		Month       int32         `json:"month" validate:"required"`
		Year        int32         `json:"year" validate:"required"`
		Description string        `json:"description"`
		UserID      string        `json:"user_id"`
		Username    string        `json:"username"`
		Branch      domain.Branch `json:"branch"`
		CreatedAt   string        `json:"created_at"`
	}

	createCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  CreateCharityMrysPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

// NewCreateCharityMrysInteractor creates new createCharityMrysInteractor with its dependencies
func NewCreateCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter CreateCharityMrysPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) CreateCharityMrysUseCase {
	return createCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a createCharityMrysInteractor) Execute(ctx context.Context, input CreateCharityMrysInput) (CreateCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	var model = domain.NewCharityMrys(
		domain.NewUUID(),
		input.Name,
		domain.Money(input.Amount),
		input.Month,
		input.Year,
		input.Description,
		input.Branch,
		time.Now(),
	)
	model.UserID = a.auth.Id
	model.Username = a.auth.Username

	model, err := a.repo.Create(ctx, model)
	if err != nil {
		return a.presenter.Output(domain.CharityMrys{}), err
	}

	return a.presenter.Output(model), nil
}

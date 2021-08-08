package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// UpdateCharityMrysUseCase input port
	UpdateCharityMrysUseCase interface {
		Execute(context.Context, UpdateCharityMrysInput, string) (UpdateCharityMrysOutput, error)
	}

	// UpdateCharityMrysInput input data
	UpdateCharityMrysInput struct {
		Id          string        `json:"id"`
		Name        string        `json:"name" validate:"required"`
		Amount      int32         `json:"amount" validate:"required"`
		Month       int32         `json:"month" validate:"required"`
		Year        int32         `json:"year" validate:"required"`
		Description string        `json:"description"`
		UserID      int32         `json:"user_id"`
		Username    string        `json:"username"`
		Branch      domain.Branch `json:"branch"`
	}

	// UpdateCharityMrysPresenter output port
	UpdateCharityMrysPresenter interface {
		Output(domain.CharityMrys) UpdateCharityMrysOutput
	}

	// UpdateCharityMrysOutput output data
	UpdateCharityMrysOutput struct {
		Id          string        `json:"id"`
		Name        string        `json:"name" validate:"required"`
		Amount      domain.Money  `json:"amount" validate:"required"`
		Month       int32         `json:"month" validate:"required"`
		Year        int32         `json:"year" validate:"required"`
		Description string        `json:"description"`
		UserID      string        `json:"user_id"`
		Username    string        `json:"username"`
		Branch      domain.Branch `json:"branch"`
		CreatedAt   string        `json:"crated_at"`
	}

	updateCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  UpdateCharityMrysPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

// NewUpdateCharityMrysInteractor updates new updateCharityMrysInteractor with its dependencies
func NewUpdateCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter UpdateCharityMrysPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) UpdateCharityMrysUseCase {
	return updateCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a updateCharityMrysInteractor) Execute(ctx context.Context, input UpdateCharityMrysInput, id string) (UpdateCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	var model = domain.NewCharityMrys(
		id,
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

	model, err := a.repo.Update(ctx, model, id)
	if err != nil {
		return a.presenter.Output(domain.CharityMrys{}), err
	}

	return a.presenter.Output(model), nil
}

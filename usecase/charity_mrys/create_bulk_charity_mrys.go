package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// CreateBulkCharityMrysInput input port
	CreateBulkCharityMrysUseCase interface {
		Execute(context.Context, CreateBulkCharityMrysInput) ([]CreateBulkCharityMrysOutput, error)
	}

	// CreateBulkCharityMrysInput input data
	CreateBulkCharityMrysInput struct {
		Id          string        `json:"id" example:"1"`
		Name        string        `json:"name" validate:"required" example:"Leo Messi"`
		Amount      int32         `json:"amount" validate:"required" example:"40000"`
		MonthFrom   int32         `json:"month_from" validate:"required" example:"2"`
		MonthTo     int32         `json:"month_to" validate:"required" example:"10"`
		Year        int32         `json:"year" validate:"required" example:"2021"`
		Branch      domain.Branch `json:"branch"`
		Description string        `json:"description" example:"description"`
	}

	// CreateBulkCharityMrysPresenter output port
	CreateBulkCharityMrysPresenter interface {
		Output([]domain.CharityMrys) []CreateBulkCharityMrysOutput
	}

	// CreateBulkCharityMrysOutput output data
	CreateBulkCharityMrysOutput struct {
		Id          string        `json:"id"`
		Name        string        `json:"name" validate:"required"`
		Amount      domain.Money  `json:"amount" validate:"required"`
		Month       int32         `json:"month" validate:"required"`
		Year        int32         `json:"year" validate:"required"`
		Description string        `json:"description"`
		Branch      domain.Branch `json:"branch"`
		CreatedAt   string        `json:"created_at"`
	}

	createBulkCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  CreateBulkCharityMrysPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateBulkCharityMrysInteractor creates new createBulkCharityMrysInteractor with its dependencies
func NewCreateBulkCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter CreateBulkCharityMrysPresenter,
	t time.Duration,
) CreateBulkCharityMrysUseCase {
	return createBulkCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a createBulkCharityMrysInteractor) Execute(ctx context.Context, input CreateBulkCharityMrysInput) ([]CreateBulkCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	data := []domain.CharityMrys{}
	for i := input.MonthFrom; i <= input.MonthTo; i++ {
		var model = domain.NewCharityMrys(
			domain.NewUUID(),
			input.Name,
			domain.Money(input.Amount),
			i,
			input.Year,
			input.Description,
			input.Branch,
			time.Now(),
		)
		data = append(data, model)
	}

	result, err := a.repo.CreateBulk(ctx, data)
	if err != nil {
		return a.presenter.Output([]domain.CharityMrys{}), err
	}

	return a.presenter.Output(result), nil
}

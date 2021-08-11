package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// UpdateReceiptLunarUseCase input port
	UpdateReceiptLunarUseCase interface {
		Execute(context.Context, UpdateReceiptLunarInput, string) (UpdateReceiptLunarOutput, error)
	}

	// UpdateReceiptLunarInput input data
	UpdateReceiptLunarInput struct {
		Id                 string                      `json:"id"`
		InternationalDate  string                      `json:"international_date" validate:"required"`
		LunarDate          string                      `json:"lunar_date" validate:"required"`
		Description        string                      `json:"description"`
		Branch             domain.Branch               `json:"branch"`
		ReceiptLunarDetail []domain.ReceiptLunarDetail `json:"receipt_lunar_detail"`
	}

	// UpdateReceiptLunarPresenter output port
	UpdateReceiptLunarPresenter interface {
		Output(domain.ReceiptLunar) UpdateReceiptLunarOutput
	}

	// UpdateReceiptLunarOutput output data
	UpdateReceiptLunarOutput struct {
		Id                 string                      `json:"id"`
		InternationalDate  time.Time                   `json:"internation_date"`
		LunarDate          string                      `json:"lunar_date"`
		Description        string                      `json:"description"`
		Total              domain.Money                `json:"total"`
		UserID             string                      `json:"user_id"`
		Username           string                      `json:"userame"`
		BranchId           string                      `json:"-"`
		Branch             domain.Branch               `json:"branch"`
		ReceiptLunarDetail []domain.ReceiptLunarDetail `json:"receipt_lunar_detail"`
		CreatedAt          string                      `json:"created_at"`
	}

	updateReceiptLunarInteractor struct {
		repo       domain.ReceiptLunarRepository
		presenter  UpdateReceiptLunarPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

// NewUpdateReceiptLunarInteractor updates new updateReceiptLunarInteractor with its dependencies
func NewUpdateReceiptLunarInteractor(
	repo domain.ReceiptLunarRepository,
	presenter UpdateReceiptLunarPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) UpdateReceiptLunarUseCase {
	return updateReceiptLunarInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

// Execute orchestrates the use case
func (a updateReceiptLunarInteractor) Execute(ctx context.Context, input UpdateReceiptLunarInput, id string) (UpdateReceiptLunarOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	internationalDate, err := time.Parse("2006-01-02", input.InternationalDate)
	if err != nil {
		return a.presenter.Output(domain.ReceiptLunar{}), err
	}

	var model = domain.NewReceiptLunar(
		id,
		internationalDate,
		input.LunarDate,
		input.Description,
		input.Branch,
		input.ReceiptLunarDetail,
		a.auth.Id,
		a.auth.Username,
		time.Now(),
	)

	model, err = a.repo.Update(ctx, model, id)
	if err != nil {
		return a.presenter.Output(domain.ReceiptLunar{}), err
	}

	return a.presenter.Output(model), nil
}

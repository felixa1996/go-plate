package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	CreateReceiptLunarUseCase interface {
		Execute(context.Context, CreateReceiptLunarInput) (CreateReceiptLunarOutput, error)
	}

	CreateReceiptLunarInput struct {
		Id                 string                      `json:"id"`
		InternationalDate  string                      `json:"international_date" validate:"required"`
		LunarDate          string                      `json:"lunar_date" validate:"required"`
		Description        string                      `json:"description"`
		Branch             domain.Branch               `json:"branch"`
		ReceiptLunarDetail []domain.ReceiptLunarDetail `json:"receipt_lunar_detail"`
	}

	CreateReceiptLunarPresenter interface {
		Output(domain.ReceiptLunar) CreateReceiptLunarOutput
	}

	CreateReceiptLunarOutput struct {
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

	createReceiptLunarInteractor struct {
		repo       domain.ReceiptLunarRepository
		presenter  CreateReceiptLunarPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

func NewCreateReceiptLunarInteractor(
	repo domain.ReceiptLunarRepository,
	presenter CreateReceiptLunarPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) CreateReceiptLunarUseCase {
	return createReceiptLunarInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

func (a createReceiptLunarInteractor) Execute(ctx context.Context, input CreateReceiptLunarInput) (CreateReceiptLunarOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	internationalDate, err := time.Parse("2006-01-02", input.InternationalDate)
	if err != nil {
		return a.presenter.Output(domain.ReceiptLunar{}), err
	}

	var model = domain.NewReceiptLunar(
		domain.NewUUID(),
		internationalDate,
		input.LunarDate,
		input.Description,
		input.Branch,
		input.ReceiptLunarDetail,
		time.Now(),
	)
	model.UserID = a.auth.Id
	model.Username = a.auth.Username

	model, err = a.repo.Create(ctx, model)
	if err != nil {
		return a.presenter.Output(domain.ReceiptLunar{}), err
	}

	return a.presenter.Output(model), nil
}

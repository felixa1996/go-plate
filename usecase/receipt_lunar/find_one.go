package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	FindReceiptLunarUseCase interface {
		Execute(context.Context, string) (FindReceiptLunarOutput, error)
	}

	FindReceiptLunarInput struct {
		ID int64 `json:"id" validate:"gt=0,required"`
	}

	FindReceiptLunarPresenter interface {
		Output(domain.ReceiptLunar) FindReceiptLunarOutput
	}

	FindReceiptLunarOutput struct {
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

	findReceiptLunarInteractor struct {
		repo       domain.ReceiptLunarRepository
		presenter  FindReceiptLunarPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

func NewFindReceiptLunarInteractor(
	repo domain.ReceiptLunarRepository,
	presenter FindReceiptLunarPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) FindReceiptLunarUseCase {
	return findReceiptLunarInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

func (a findReceiptLunarInteractor) Execute(ctx context.Context, ID string) (FindReceiptLunarOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	model, err := a.repo.FindByID(ctx, ID)
	if err != nil {
		return a.presenter.Output(domain.ReceiptLunar{}), err
	}

	return a.presenter.Output(model), nil
}

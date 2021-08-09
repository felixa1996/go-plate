package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	FindPaginationReceiptLunarUseCase interface {
		Execute(context.Context, int, int, string, string) (FindPaginationReceiptLunarOutput, error)
	}

	FindPaginationReceiptLunarPresenter interface {
		Output(domain.ReceiptLunarPagination) FindPaginationReceiptLunarOutput
	}

	FindPaginationReceiptLunarOutput struct {
		Data []FindPaginationReceiptLunarOutputData `json:"data"`
		Meta domain.MetaPagination                  `json:"meta"`
	}

	findPaginationReceiptLunarInteractor struct {
		repo       domain.ReceiptLunarRepository
		presenter  FindPaginationReceiptLunarPresenter
		auth       *domain.UserJwt
		ctxTimeout time.Duration
	}
)

func NewFindPaginationReceiptLunarInteractor(
	repo domain.ReceiptLunarRepository,
	presenter FindPaginationReceiptLunarPresenter,
	auth *domain.UserJwt,
	t time.Duration,
) FindPaginationReceiptLunarUseCase {
	return findPaginationReceiptLunarInteractor{
		repo:       repo,
		presenter:  presenter,
		auth:       auth,
		ctxTimeout: t,
	}
}

func (a findPaginationReceiptLunarInteractor) Execute(ctx context.Context, currentPage int, perPage int, sort string, search string) (FindPaginationReceiptLunarOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	result, err := a.repo.FindPagination(ctx, currentPage, perPage, sort, search)
	if err != nil {
		return a.presenter.Output(result), err
	}

	return a.presenter.Output(result), nil
}

type FindPaginationReceiptLunarOutputData struct {
	Id                string        `json:"id"`
	InternationalDate time.Time     `json:"internation_date"`
	LunarDate         string        `json:"lunar_date"`
	Description       string        `json:"description"`
	Total             domain.Money  `json:"total"`
	UserID            string        `json:"user_id"`
	Username          string        `json:"userame"`
	BranchId          string        `json:"-"`
	Branch            domain.Branch `json:"branch"`
	CreatedAt         string        `json:"created_at"`
}

package usecase

import (
	"context"
	"time"

	"github.com/felixa1996/go-plate/domain"
)

type (
	// FindPaginationCharityMrysUseCase input port
	FindPaginationCharityMrysUseCase interface {
		Execute(context.Context, int, int, string, string) (FindPaginationCharityMrysOutput, error)
	}

	FindPaginationCharityMrysPresenter interface {
		Output(domain.CharityMrysPagination) FindPaginationCharityMrysOutput
	}

	FindPaginationCharityMrysOutput struct {
		Data []FindPaginationCharityMrysOutputData `json:"data"`
		Meta domain.MetaPagination                 `json:"meta"`
	}

	findPaginationCharityMrysInteractor struct {
		repo       domain.CharityMrysRepository
		presenter  FindPaginationCharityMrysPresenter
		ctxTimeout time.Duration
	}
)

func NewFindPaginationCharityMrysInteractor(
	repo domain.CharityMrysRepository,
	presenter FindPaginationCharityMrysPresenter,
	t time.Duration,
) FindPaginationCharityMrysUseCase {
	return findPaginationCharityMrysInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: t,
	}
}

func (a findPaginationCharityMrysInteractor) Execute(ctx context.Context, currentPage int, perPage int, sort string, search string) (FindPaginationCharityMrysOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, a.ctxTimeout)
	defer cancel()

	result, err := a.repo.FindPagination(ctx, currentPage, perPage, sort, search)
	if err != nil {
		return a.presenter.Output(result), err
	}

	return a.presenter.Output(result), nil
}

// List of response
type FindPaginationCharityMrysOutputData struct {
	Id          string        `json:"id"`
	Name        string        `json:"name"`
	Amount      float64       `json:"Amount"`
	Month       int32         `json:"Month"`
	Year        int32         `json:"Year"`
	Description string        `json:"Description"`
	UserID      string        `json:"user_id"`
	Username    string        `json:"username"`
	Branch      domain.Branch `json:"branch"`
	CreatedAt   time.Time     `json:"CreatedAt"`
}

package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrBranchNotFound = errors.New("Branch not found")
)

type Branch struct {
	Id          string    `json:"id" example:"1"`
	Code        string    `json:"code" example:"PTK"`
	Name        string    `json:"name" example:"Pontianak"`
	Address     string    `json:"address" example:"Pontianak"`
	Description string    `json:"Description" example:"description"`
	CreatedAt   time.Time `json:"CreatedAt" example:"2019-11-09T21:21:46+00:00"`
}

type BranchPagination struct {
	Data []Branch       `json:"data"`
	Meta MetaPagination `json:"meta"`
}
type BranchAll struct {
	Data []Branch `json:"data"`
}

type (
	BranchRepository interface {
		Create(context.Context, Branch) (Branch, error)
		Update(context.Context, Branch, string) (Branch, error)
		FindPagination(context.Context, int, int, string, string) (BranchPagination, error)
		FindAll(context.Context) (BranchAll, error)
		FindByID(context.Context, string) (Branch, error)
		DeleteByID(context.Context, string) (bool, error)
	}
)

func DeleteBranch(success bool) bool {
	return success
}

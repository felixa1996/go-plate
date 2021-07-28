package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCharityMrysNotFound = errors.New("CharityMrys not found")
)

type CharityMrys struct {
	Id          string    `json:"id" example:"1"`
	Name        string    `json:"name" example:"Leo Messi"`
	Amount      Money     `json:"Amount" example:"40000"`
	Month       int32     `json:"Month" example:"1"`
	Year        int32     `json:"Year" example:"2021"`
	Description string    `json:"Description" example:"description"`
	CreatedAt   time.Time `json:"CreatedAt" example:"2019-11-09T21:21:46+00:00"`
}

type CharityMrysPagination struct {
	Data []CharityMrys  `json:"data"`
	Meta MetaPagination `json:"meta"`
}
type CharityMrysAll struct {
	Data []CharityMrys `json:"data"`
}

type (
	CharityMrysRepository interface {
		CreateBulk(context.Context, []CharityMrys) ([]CharityMrys, error)
		Create(context.Context, CharityMrys) (CharityMrys, error)
		Update(context.Context, CharityMrys, string) (CharityMrys, error)
		FindPagination(context.Context, int, int, string, string) (CharityMrysPagination, error)
		FindAll(context.Context) (CharityMrysAll, error)
		FindByID(context.Context, string) (CharityMrys, error)
		DeleteByID(context.Context, string) (bool, error)
	}
)

func NewCharityMrys(ID string, Name string, Amount Money, Month int32, Year int32, Description string, createdAt time.Time) CharityMrys {
	return CharityMrys{
		Id:          ID,
		Name:        Name,
		Amount:      Amount,
		Month:       Month,
		Year:        Year,
		Description: Description,
		CreatedAt:   createdAt,
	}
}

func DeleteCharityMrys(success bool) bool {
	return success
}

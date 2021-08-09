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
	Id          string    `gorm:"primaryKey" json:"id" example:"1"`
	Name        string    `json:"name" example:"Leo Messi"`
	Amount      Money     `json:"Amount" example:"40000"`
	Month       int32     `json:"Month" example:"1"`
	Year        int32     `json:"Year" example:"2021"`
	Description string    `json:"Description" example:"description"`
	UserID      string    `json:"user_id" example:"1"`
	Username    string    `json:"Userame" example:"usename"`
	BranchId    string    `json:"-"`
	Branch      Branch    `json:"branch" gorm:"foreignKey:BranchId;branch:ID"`
	CreatedAt   time.Time `json:"created_at" example:"2019-11-09T21:21:46+00:00"`
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
		FindAll(context.Context, *UserJwt) (CharityMrysAll, error)
		FindByID(context.Context, string) (CharityMrys, error)
		DeleteByID(context.Context, string) (bool, error)
	}
)

func NewCharityMrys(ID string, Name string, Amount Money, Month int32, Year int32, Description string, Branch Branch, createdAt time.Time) CharityMrys {
	return CharityMrys{
		Id:          ID,
		Name:        Name,
		Amount:      Amount,
		Month:       Month,
		Year:        Year,
		Description: Description,
		Branch:      Branch,
		CreatedAt:   createdAt,
	}
}

func DeleteCharityMrys(success bool) bool {
	return success
}

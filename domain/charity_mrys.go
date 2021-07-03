package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrCharityMrysNotFound = errors.New("CharityMrys not found")
)

type CharityMrysID string

func (a CharityMrysID) String() string {
	return string(a)
}

type CharityMrys struct {
	Id          CharityMrysID `json:"id" example:"1"`
	Name        string        `json:"name" example:"Leo Messi"`
	Amount      Money         `json:"Amount" example:"40000"`
	Month       int32         `json:"Month" example:"1"`
	Year        int32         `json:"Year" example:"2021"`
	Description string        `json:"Description" example:"description"`
	CreatedAt   time.Time     `json:"CreatedAt" example:"2019-11-09T21:21:46+00:00"`
}

type (
	CharityMrysRepository interface {
		Create(context.Context, CharityMrys) (CharityMrys, error)
		Update(context.Context, CharityMrysID, CharityMrys) (CharityMrys, error)
		FindAll(context.Context) ([]CharityMrys, error)
		FindByID(context.Context, CharityMrysID) (CharityMrys, error)
		DeleteByID(context.Context, CharityMrysID) (bool, error)
	}
)

func NewCharityMrys(ID CharityMrysID, Name string, Amount Money, Month int32, Year int32, Description string, createdAt time.Time) CharityMrys {
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

func (a CharityMrys) ID() CharityMrysID {
	return a.Id
}

func DeleteCharityMrys(success bool) bool {
	return success
}

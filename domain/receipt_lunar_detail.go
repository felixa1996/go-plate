package domain

import (
	"context"
	"errors"
	"time"
)

var (
	ErrReceiptLunarDetailNotFound = errors.New("ReceiptLunarDetail not found")
)

type ReceiptLunarDetail struct {
	Id             string       `gorm:"primaryKey" json:"id" example:"1"`
	Name           string       `json:"name"`
	Description    string       `json:"description"`
	Total          Money        `json:"total" example:"40000"`
	ReceiptLunarId string       `json:"-"`
	ReceiptLunar   ReceiptLunar `json:"receipt_lunar" gorm:"foreignKey:ReceiptLunarId;receipt_lunar:ID"`
	CreatedAt      time.Time    `json:"created_at" example:"2019-11-09T21:21:46+00:00"`
}

type ReceiptLunarDetailAll struct {
	Data []ReceiptLunarDetail `json:"data"`
}

type (
	ReceiptLunarDetailRepository interface {
		CreateBulk(context.Context, []ReceiptLunarDetail) ([]ReceiptLunarDetail, error)
		DeleteByReceiptLunarId(context.Context, string) (bool, error)
	}
)

func NewReceiptLunarDetail(Id string, Name string, Description string, Total Money, ReceiptLunar ReceiptLunar, createdAt time.Time) ReceiptLunarDetail {
	return ReceiptLunarDetail{
		Id:           Id,
		Name:         Name,
		Description:  Description,
		Total:        Total,
		ReceiptLunar: ReceiptLunar,
		CreatedAt:    createdAt,
	}
}

func DeleteReceiptLunarDetail(success bool) bool {
	return success
}

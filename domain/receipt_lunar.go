package domain

import (
	"context"
	"errors"
	"time"

	gouuid "github.com/satori/go.uuid"
)

var (
	ErrReceiptLunarNotFound = errors.New("ReceiptLunar not found")
)

type ReceiptLunar struct {
	Id                 string               `gorm:"primaryKey" json:"id" example:"1"`
	InternationalDate  time.Time            `json:"internation_date" example:"2021-08-04"`
	LunarDate          string               `json:"lunar_date" example:"2021-08-04"`
	Description        string               `json:"description"`
	Total              Money                `json:"total" example:"40000"`
	UserID             string               `json:"user_id" example:"1"`
	Username           string               `json:"username" example:"username"`
	BranchId           string               `json:"-"`
	Branch             Branch               `json:"branch" gorm:"foreignKey:BranchId;branch:ID"`
	ReceiptLunarDetail []ReceiptLunarDetail `json:"receipt_lunar_detail"`
	CreatedAt          time.Time            `json:"created_at" example:"2019-11-09T21:21:46+00:00"`
}

type ReceiptLunarPagination struct {
	Data []ReceiptLunar `json:"data"`
	Meta MetaPagination `json:"meta"`
}
type ReceiptLunarAll struct {
	Data []ReceiptLunar `json:"data"`
}

type (
	ReceiptLunarRepository interface {
		Create(context.Context, ReceiptLunar) (ReceiptLunar, error)
		Update(context.Context, ReceiptLunar, string) (ReceiptLunar, error)
		FindPagination(context.Context, int, int, string, string) (ReceiptLunarPagination, error)
		FindByID(context.Context, string) (ReceiptLunar, error)
		DeleteByID(context.Context, string) (bool, error)
	}
)

func NewReceiptLunar(Id string, InternationalDate time.Time, LunarDate string, Description string, Branch Branch, ReceiptLunarDetail []ReceiptLunarDetail, UserId string, Username string, CreatedAt time.Time) ReceiptLunar {
	return ReceiptLunar{
		Id:                 Id,
		InternationalDate:  InternationalDate,
		LunarDate:          LunarDate,
		Description:        Description,
		Branch:             Branch,
		ReceiptLunarDetail: ReceiptLunarDetail,
		UserID:             UserId,
		Username:           Username,
		CreatedAt:          CreatedAt,
	}
}

func (r ReceiptLunar) ProcessDetail(ReceiptLunarDetailList []ReceiptLunarDetail) ReceiptLunar {
	result := float64(0)
	detail := []ReceiptLunarDetail{}
	for _, v := range ReceiptLunarDetailList {
		v.Id = gouuid.NewV4().String()
		detail = append(detail, v)
		result += float64(v.Total)
	}
	r.ReceiptLunarDetail = detail
	r.Total = Money(result)
	return r
}

func DeleteReceiptLunar(success bool) bool {
	return success
}

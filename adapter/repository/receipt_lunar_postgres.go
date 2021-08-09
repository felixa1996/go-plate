package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/felixa1996/go-plate/domain"
	utils "github.com/felixa1996/go-plate/infrastructure/utils"
	"github.com/pkg/errors"
)

type ReceiptLunarSQL struct {
	db SQL
}

func NewReceiptLunarSQL(db SQL) ReceiptLunarSQL {
	return ReceiptLunarSQL{
		db: db,
	}
}

func (a ReceiptLunarSQL) Create(ctx context.Context, ReceiptLunar domain.ReceiptLunar) (domain.ReceiptLunar, error) {
	if len(ReceiptLunar.ReceiptLunarDetail) == 0 {
		return domain.ReceiptLunar{}, errors.Wrap(nil, "error empty ReceiptLunarDetail")
	}

	ReceiptLunar = ReceiptLunar.ProcessDetail(ReceiptLunar.ReceiptLunarDetail)

	if err := a.db.GetDBGorm(ctx).Create(&ReceiptLunar).Error; err != nil {
		return domain.ReceiptLunar{}, errors.Wrap(err, "error creating ReceiptLunar")
	}

	return ReceiptLunar, nil
}

func (a ReceiptLunarSQL) Update(ctx context.Context, ReceiptLunar domain.ReceiptLunar, Id string) (domain.ReceiptLunar, error) {
	if len(ReceiptLunar.ReceiptLunarDetail) == 0 {
		return domain.ReceiptLunar{}, errors.Wrap(nil, "error empty ReceiptLunarDetail")
	}

	a.DeleteByReceiptLunarId(ctx, Id)
	ReceiptLunar = ReceiptLunar.ProcessDetail(ReceiptLunar.ReceiptLunarDetail)

	whereQuery := fmt.Sprintf("id = '%s'", Id)
	if err := a.db.UpdatePG(ctx, &ReceiptLunar, whereQuery); err != nil {
		return domain.ReceiptLunar{}, errors.Wrap(err, "error updating ReceiptLunar")
	}

	return ReceiptLunar, nil
}

func (a ReceiptLunarSQL) FindPagination(ctx context.Context, currentPage int, perPage int, sort string, search string) (domain.ReceiptLunarPagination, error) {
	db := a.db.GetDBGorm(ctx)
	var total int64
	err := db.Model(&domain.ReceiptLunar{}).Count(&total)
	if err.Error != nil {
		return domain.ReceiptLunarPagination{}, errors.Wrap(err.Error, "error get total ReceiptLunars")
	}

	pagination := utils.Pagination{
		PerPage:     perPage,
		CurrentPage: currentPage,
		Total:       total,
		Sort:        sort,
	}
	meta := pagination.ToMeta()

	q := db.Model(&domain.ReceiptLunar{}).
		Limit(perPage).
		Order(meta.Sort).
		Offset(meta.Offset)

	if len(search) > 0 {
		q.Where("LOWER(description) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	var list []domain.ReceiptLunar

	err = q.Find(&list)
	if err.Error != nil {
		return domain.ReceiptLunarPagination{}, errors.Wrap(err.Error, "error listing pagination ReceiptLunars")
	}

	return domain.ReceiptLunarPagination{
		Data: list,
		Meta: meta,
	}, nil
}

func (a ReceiptLunarSQL) FindByID(ctx context.Context, Id string) (domain.ReceiptLunar, error) {
	var one domain.ReceiptLunar

	db := a.db.GetDBGorm(ctx)
	res := db.Model(&domain.ReceiptLunar{}).Where("id = ?", Id).First(&one)

	if res.Error != nil {
		return domain.ReceiptLunar{}, domain.ErrReceiptLunarNotFound
	}

	detail, err := a.FindAllByReceiptLunarId(ctx, Id)
	if err != nil {
		return domain.ReceiptLunar{}, err
	}
	one.ReceiptLunarDetail = detail

	return one, nil
}

func (a ReceiptLunarSQL) DeleteByID(ctx context.Context, Id string) (bool, error) {
	var result domain.ReceiptLunar
	db := a.db.GetDBGorm(ctx)
	res := db.Where("id = ?", Id).Delete(result)
	if res.Error != nil {
		return false, domain.ErrReceiptLunarNotFound
	}
	return true, nil
}

func (a ReceiptLunarSQL) DeleteByReceiptLunarId(ctx context.Context, Id string) (bool, error) {
	var result domain.ReceiptLunarDetail
	db := a.db.GetDBGorm(ctx)
	res := db.Where("receipt_lunar_id = ?", Id).Delete(result)
	if res.Error != nil {
		return false, domain.ErrReceiptLunarDetailNotFound
	}
	return true, nil
}

func (a ReceiptLunarSQL) FindAllByReceiptLunarId(ctx context.Context, Id string) ([]domain.ReceiptLunarDetail, error) {
	var list []domain.ReceiptLunarDetail

	db := a.db.GetDBGorm(ctx)
	res := db.Model(&domain.ReceiptLunarDetail{}).
		Where("receipt_lunar_id = ?", Id).
		Order("name ASC").
		Find(&list)

	if res.Error != nil {
		return []domain.ReceiptLunarDetail{}, errors.Wrap(res.Error, "error listing all ReceiptLunarDetail")
	}

	return list, nil
}

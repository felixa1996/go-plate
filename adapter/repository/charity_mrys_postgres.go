package repository

import (
	"context"
	"fmt"
	"strings"

	"github.com/felixa1996/go-plate/domain"
	utils "github.com/felixa1996/go-plate/infrastructure/utils"
	"github.com/pkg/errors"
)

type CharityMrysSQL struct {
	db SQL
}

func NewCharityMrysSQL(db SQL) CharityMrysSQL {
	return CharityMrysSQL{
		db: db,
	}
}

func (a CharityMrysSQL) CreateBulk(ctx context.Context, CharityMrys []domain.CharityMrys) ([]domain.CharityMrys, error) {
	if err := a.db.GetDBGorm(ctx).Create(&CharityMrys).Error; err != nil {
		return []domain.CharityMrys{}, errors.Wrap(err, "error creating CharityMrysBulk")
	}

	return CharityMrys, nil
}

func (a CharityMrysSQL) Create(ctx context.Context, CharityMrys domain.CharityMrys) (domain.CharityMrys, error) {
	if err := a.db.GetDBGorm(ctx).Create(&CharityMrys).Error; err != nil {
		return domain.CharityMrys{}, errors.Wrap(err, "error creating CharityMrys")
	}

	return CharityMrys, nil
}

func (a CharityMrysSQL) Update(ctx context.Context, CharityMrys domain.CharityMrys, ID string) (domain.CharityMrys, error) {
	/* There're 2 ways to perform update data using modelContext for simple one
	 * If query is complex we're using getDBPG to get the handler and write down the query inside repostiory
	 */
	whereQuery := fmt.Sprintf("id = '%s'", ID)

	// db := a.db.GetDBPG(ctx)
	// _, err := db.Model(&CharityMrys).Where(whereQuery).Update()
	// if err != nil {
	// 	return domain.CharityMrys{}, errors.Wrap(err, "error updating CharityMrys")
	// }

	if err := a.db.UpdatePG(ctx, &CharityMrys, whereQuery); err != nil {
		return domain.CharityMrys{}, errors.Wrap(err, "error updating CharityMrys")
	}

	return CharityMrys, nil
}

func (a CharityMrysSQL) FindPagination(ctx context.Context, currentPage int, perPage int, sort string, search string) (domain.CharityMrysPagination, error) {
	db := a.db.GetDBGorm(ctx)
	var total int64
	err := db.Model(&domain.CharityMrys{}).Count(&total)
	if err.Error != nil {
		return domain.CharityMrysPagination{}, errors.Wrap(err.Error, "error get total CharityMryss")
	}

	pagination := utils.Pagination{
		PerPage:     perPage,
		CurrentPage: currentPage,
		Total:       total,
		Sort:        sort,
		// SortList: map[string]string{
		// 	"name2": "name",
		// },
	}
	meta := pagination.ToMeta()

	q := db.Model(&domain.CharityMrys{}).
		Preload("Branch").
		Limit(perPage).
		Order(meta.Sort).
		Offset(meta.Offset)

	if len(search) > 0 {
		q.Where("LOWER(charity_mrys.name) LIKE ?", "%"+strings.ToLower(search)+"%")
	}

	var list []domain.CharityMrys

	err = q.Find(&list)
	if err.Error != nil {
		return domain.CharityMrysPagination{}, errors.Wrap(err.Error, "error listing pagination CharityMryss")
	}

	return domain.CharityMrysPagination{
		Data: list,
		Meta: meta,
	}, nil
}

func (a CharityMrysSQL) FindAll(ctx context.Context, auth *domain.UserJwt) (domain.CharityMrysAll, error) {
	var list []domain.CharityMrys

	db := a.db.GetDBGorm(ctx)
	res := db.Model(&domain.CharityMrys{}).
		Preload("Branch").
		Order("charity_mrys.year ASC,charity_mrys.month ASC").
		Find(&list)

	if res.Error != nil {
		return domain.CharityMrysAll{}, errors.Wrap(res.Error, "error listing all CharityMryss")
	}

	return domain.CharityMrysAll{
		Data: list,
	}, nil
}

func (a CharityMrysSQL) FindByID(ctx context.Context, ID string) (domain.CharityMrys, error) {
	var one domain.CharityMrys

	db := a.db.GetDBGorm(ctx)
	res := db.Model(&domain.CharityMrys{}).
		Preload("Branch").
		Where("charity_mrys.id = ?", ID).First(&one)

	if res.Error != nil {
		return domain.CharityMrys{}, domain.ErrCharityMrysNotFound
	}
	return one, nil
}

func (a CharityMrysSQL) DeleteByID(ctx context.Context, Id string) (bool, error) {

	var result domain.CharityMrys
	db := a.db.GetDBGorm(ctx)
	res := db.Where("id = ?", Id).Delete(result)
	if res.Error != nil {
		return false, domain.ErrCharityMrysNotFound
	}
	return true, nil
}

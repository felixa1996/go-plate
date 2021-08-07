package repository

import (
	"context"
	"database/sql"
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
	if err := a.db.InsertPG(ctx, "name", &CharityMrys); err != nil {
		return []domain.CharityMrys{}, errors.Wrap(err, "error creating CharityMrys")
	}
	return CharityMrys, nil
}

func (a CharityMrysSQL) Create(ctx context.Context, CharityMrys domain.CharityMrys) (domain.CharityMrys, error) {
	if err := a.db.GetDBGorm(ctx).Set("gorm:save_associations", false).Create(&CharityMrys).Error; err != nil {
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
		Preload("Branch").Limit(perPage).Order(meta.Sort).Offset(meta.Offset)
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
	var query = "SELECT * FROM charity_mrys"

	var list []domain.CharityMrys

	_, err := a.db.QueryContextPG(ctx, &list, query)
	if err != nil {
		return domain.CharityMrysAll{}, errors.Wrap(err, "error listing CharityMryss")
	}

	return domain.CharityMrysAll{
		Data: list,
	}, nil
}

func (a CharityMrysSQL) FindByID(ctx context.Context, ID string) (domain.CharityMrys, error) {
	tx, ok := ctx.Value("TransactionContextKey").(Tx)
	if !ok {
		var err error
		tx, err = a.db.BeginTx(ctx)
		if err != nil {
			return domain.CharityMrys{}, errors.Wrap(err, "error find CharityMrys by id")
		}
	}

	var one domain.CharityMrys

	query := "SELECT * FROM charity_mrys WHERE id = ? LIMIT 1"

	_, err := tx.QueryRowContextPG(ctx, &one, query, ID)
	if err != nil {
		return domain.CharityMrys{}, domain.ErrCharityMrysNotFound
	}
	return one, nil
}

func (a CharityMrysSQL) DeleteByID(ctx context.Context, ID string) (bool, error) {

	var (
		query = "DELETE FROM charity_mrys WHERE id = ?"
	)

	var result domain.CharityMrys

	err := a.db.ExecuteContextPG(ctx, &result, query, ID)
	switch {
	case err == sql.ErrNoRows:
		return false, domain.ErrCharityMrysNotFound
	default:
		return domain.DeleteCharityMrys(true), err
	}
}

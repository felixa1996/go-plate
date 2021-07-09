package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/felixa1996/go-plate/domain"
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

func (a CharityMrysSQL) Create(ctx context.Context, CharityMrys domain.CharityMrys) (domain.CharityMrys, error) {
	// var query = `
	// 	INSERT INTO
	// 		charity_mrys (id, name, amount, year, month, description, created_at)
	// 	VALUES
	// 		($1, $2, $3, $4, $5, $6, $7)
	// `

	if err := a.db.InsertPG(ctx, &CharityMrys, "name"); err != nil {
		return domain.CharityMrys{}, errors.Wrap(err, "error creating CharityMrys")
	}
	fmt.Printf(CharityMrys.Id)

	return CharityMrys, nil
}

func (a CharityMrysSQL) Update(ctx context.Context, ID string, charityMrys domain.CharityMrys) (domain.CharityMrys, error) {
	tx, ok := ctx.Value("TransactionContextKey").(Tx)
	if !ok {
		var err error
		tx, err = a.db.BeginTx(ctx)
		if err != nil {
			return domain.CharityMrys{}, errors.Wrap(err, "error updating CharityMrys transaction")
		}
	}

	query := "UPDATE charity_mrys SET name = $2, amount = $3, year = $4, month = $5, description = $6 WHERE id = $1"

	if err := tx.ExecuteContext(ctx, query, charityMrys, ID); err != nil {
		return charityMrys, errors.Wrap(err, "error updating CharityMrys query")
	}

	return charityMrys, nil
}

func (a CharityMrysSQL) FindAll(ctx context.Context) ([]domain.CharityMrys, error) {
	var query = "SELECT * FROM charity_mrys"

	var list []domain.CharityMrys

	_, err := a.db.QueryContextPG(ctx, &list, query)
	if err != nil {
		return []domain.CharityMrys{}, errors.Wrap(err, "error listing CharityMryss")
	}

	return list, nil
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
		//todo check
		return domain.DeleteCharityMrys(true), err
	}
}

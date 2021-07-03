package repository

import (
	"context"
	"database/sql"
	"time"

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
	var query = `
		INSERT INTO 
			charity_mrys (id, name, amount, year, month, description, created_at)
		VALUES 
			($1, $2, $3, $4, $5, $6, $7)
	`

	if err := a.db.ExecuteContext(
		ctx,
		query,
		CharityMrys.ID(),
		CharityMrys.Name,
		CharityMrys.Amount,
		CharityMrys.Year,
		CharityMrys.Month,
		CharityMrys.Description,
		CharityMrys.CreatedAt,
	); err != nil {
		return domain.CharityMrys{}, errors.Wrap(err, "error creating CharityMrys")
	}

	return CharityMrys, nil
}

func (a CharityMrysSQL) Update(ctx context.Context, ID domain.CharityMrysID, charityMrys domain.CharityMrys) (domain.CharityMrys, error) {
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

	rows, err := a.db.QueryContext(ctx, query)
	if err != nil {
		return []domain.CharityMrys{}, errors.Wrap(err, "error listing CharityMryss")
	}

	var CharityMryss = make([]domain.CharityMrys, 0)
	for rows.Next() {
		var (
			ID          string
			Name        string
			Amount      int64
			Month       int32
			Year        int32
			Description string
			createdAt   time.Time
		)

		if err = rows.Scan(&ID, &Name, &Amount, &Month, &Year, &Description, &createdAt); err != nil {
			return []domain.CharityMrys{}, errors.Wrap(err, "error listing CharityMryss")
		}

		CharityMryss = append(CharityMryss, domain.NewCharityMrys(
			domain.CharityMrysID(ID),
			Name,
			domain.Money(Amount),
			Month,
			Year,
			Description,
			createdAt,
		))
	}
	defer rows.Close()

	if err = rows.Err(); err != nil {
		return []domain.CharityMrys{}, err
	}

	return CharityMryss, nil
}

func (a CharityMrysSQL) FindByID(ctx context.Context, ID domain.CharityMrysID) (domain.CharityMrys, error) {
	tx, ok := ctx.Value("TransactionContextKey").(Tx)
	if !ok {
		var err error
		tx, err = a.db.BeginTx(ctx)
		if err != nil {
			return domain.CharityMrys{}, errors.Wrap(err, "error find CharityMrys by id")
		}
	}

	var (
		query       = "SELECT * FROM charity_mrys WHERE id = $1 LIMIT 1 FOR NO KEY UPDATE"
		id          string
		name        string
		amount      int32
		month       int32
		year        int32
		description string
		createdAt   time.Time
	)

	err := tx.QueryRowContext(ctx, query, ID).Scan(&id, &name, &amount, &month, &year, &description, &createdAt)
	switch {
	case err == sql.ErrNoRows:
		return domain.CharityMrys{}, domain.ErrCharityMrysNotFound
	default:
		return domain.NewCharityMrys(
			domain.CharityMrysID(id),
			name,
			domain.Money(amount),
			month,
			year,
			description,
			createdAt,
		), err
	}
}

func (a CharityMrysSQL) DeleteByID(ctx context.Context, ID domain.CharityMrysID) (bool, error) {

	var (
		query = "DELETE FROM charity_mrys WHERE id = $1"
	)

	err := a.db.ExecuteContext(ctx, query, ID)
	switch {
	case err == sql.ErrNoRows:
		return false, domain.ErrCharityMrysNotFound
	default:
		//todo check
		return domain.DeleteCharityMrys(true), err
	}
}

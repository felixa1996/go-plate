package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/felixa1996/go-plate/adapter/repository"
	"github.com/felixa1996/go-plate/domain"
	"github.com/go-pg/pg/extra/pgdebug"
	"github.com/go-pg/pg/v10"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type postgresHandler struct {
	db     *sql.DB
	dbPG   *pg.DB
	dbGorm *gorm.DB
}

func NewPostgresHandler(c *config) (*postgresHandler, error) {
	var ds = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
	)

	fmt.Println(ds)
	db, err := sql.Open(c.driver, ds)
	if err != nil {
		return &postgresHandler{}, err
	}

	err = db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	var dsPG = fmt.Sprintf(
		"%s://%s:%s@%s:%s/%s?sslmode=disable",
		c.driver,
		c.user,
		c.password,
		c.host,
		c.port,
		c.database,
	)

	fmt.Println(dsPG)
	opt, err := pg.ParseURL(dsPG)
	if err != nil {
		return &postgresHandler{}, err
	}

	dbPG := pg.Connect(opt)
	if err != nil {
		log.Fatalln(err)
	}
	dbPG.AddQueryHook(pgdebug.DebugHook{
		// Print all queries.
		Verbose: true,
	})

	var dsGorm = fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		c.host,
		c.port,
		c.user,
		c.database,
		c.password,
	)

	dbGorm, err := gorm.Open(postgres.Open(dsGorm), &gorm.Config{
		PrepareStmt:            true,
		SkipDefaultTransaction: false,
		FullSaveAssociations:   false,
	})
	if err != nil {
		log.Fatalln(err)
	}

	dbGorm.AutoMigrate(&domain.Branch{})
	dbGorm.AutoMigrate(&domain.CharityMrys{})
	dbGorm.AutoMigrate(&domain.ReceiptLunar{})
	dbGorm.AutoMigrate(&domain.ReceiptLunarDetail{})

	return &postgresHandler{db: db, dbPG: dbPG, dbGorm: dbGorm}, nil
}

func (p postgresHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
	tx, err := p.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return postgresTx{}, err
	}

	txPg, err := p.dbPG.BeginContext(ctx)
	if err != nil {
		return postgresTx{}, err
	}

	return newPostgresTx(tx, txPg), nil
}

func (p postgresHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresHandler) ExecuteContextPG(ctx context.Context, model interface{}, query string, args ...interface{}) error {
	_, err := p.dbPG.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresHandler) InsertPG(ctx context.Context, returning string, model ...interface{}) error {
	_, err := p.dbPG.ModelContext(ctx, model).Insert()
	if err != nil {
		return err
	}

	return nil
}

func (p postgresHandler) UpdatePG(ctx context.Context, model interface{}, where string) error {
	_, err := p.dbPG.ModelContext(ctx, model).Where(where).Update()
	if err != nil {
		return err
	}

	return nil
}

func (p postgresHandler) GetDBPG(ctx context.Context) *pg.DB {
	return p.dbPG
}

func (p postgresHandler) GetDBGorm(ctx context.Context) *gorm.DB {
	return p.dbGorm
}

func (p postgresHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := p.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newPostgresRows(rows)

	return row, nil
}

func (p postgresHandler) QueryContextPG(ctx context.Context, model interface{}, query string, args ...interface{}) (pg.Result, error) {
	rows, err := p.dbPG.QueryContext(ctx, model, query, args...)
	if err != nil {
		return nil, err
	}
	return rows, nil
}

func (p postgresHandler) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := p.db.QueryRowContext(ctx, query, args...)

	return newPostgresRow(row)
}

func (p postgresHandler) QueryRowContextPG(ctx context.Context, model interface{}, query string, args ...interface{}) (pg.Result, error) {
	row, err := p.dbPG.QueryOneContext(ctx, model, query, args...)
	if err != nil {
		return nil, err
	}

	return row, nil
}

type postgresRow struct {
	row *sql.Row
}

func newPostgresRow(row *sql.Row) postgresRow {
	return postgresRow{row: row}
}

func (pr postgresRow) Scan(dest ...interface{}) error {
	if err := pr.row.Scan(dest...); err != nil {
		return err
	}

	return nil
}

type postgresRows struct {
	rows *sql.Rows
}

func newPostgresRows(rows *sql.Rows) postgresRows {
	return postgresRows{rows: rows}
}

func (pr postgresRows) Scan(dest ...interface{}) error {
	if err := pr.rows.Scan(dest...); err != nil {
		return err
	}

	return nil
}

func (pr postgresRows) Next() bool {
	return pr.rows.Next()
}

func (pr postgresRows) Err() error {
	return pr.rows.Err()
}

func (pr postgresRows) Close() error {
	return pr.rows.Close()
}

type postgresTx struct {
	tx   *sql.Tx
	txPg *pg.Tx
}

func newPostgresTx(tx *sql.Tx, txPg *pg.Tx) postgresTx {
	return postgresTx{tx: tx, txPg: txPg}
}

func (p postgresTx) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
	_, err := p.tx.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresTx) ExecuteContextPG(ctx context.Context, model interface{}, query string, args ...interface{}) error {
	_, err := p.txPg.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (p postgresTx) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
	rows, err := p.tx.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	row := newPostgresRows(rows)

	return row, nil
}

func (p postgresTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
	row := p.tx.QueryRowContext(ctx, query, args...)

	return newPostgresRow(row)
}

func (p postgresTx) QueryRowContextPG(ctx context.Context, model interface{}, query string, args ...interface{}) (pg.Result, error) {
	row, err := p.txPg.QueryOneContext(ctx, model, query, args...)
	if err != nil {
		return nil, err
	}

	return row, nil
}

func (p postgresTx) Commit() error {
	return p.tx.Commit()
}

func (p postgresTx) Rollback() error {
	return p.tx.Rollback()
}

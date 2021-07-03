package database

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	"github.com/felixa1996/go-plate/adapter/repository"

// 	_ "github.com/lib/pq"

// 	"github.com/go-pg/pg/v10"
// )

// type pgHandler struct {
// 	db *pg.DB
// }

// func NewPgHandler(c *config) (*pgHandler, error) {
// 	var ds = fmt.Sprintf(
// 		"%s://%s:%s@%s:%s/%s",
// 		c.host,
// 		c.user,
// 		c.password,
// 		c.port,
// 		c.database,
// 	)

// 	fmt.Println(ds)
// 	opt, err := pg.ParseURL(ds)
// 	if err != nil {
// 		return &pgHandler{}, err
// 	}

// 	db := pg.Connect(opt)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}

// 	return &pgHandler{db: db}, nil
// }

// func (p pgHandler) BeginTx(ctx context.Context) (repository.Tx, error) {
// 	tx, err := p.db.BeginContext(ctx)
// 	if err != nil {
// 		return pgsTx{}, err
// 	}

// 	return newPgsTx(tx), nil
// }

// func (p pgHandler) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
// 	_, err := p.db.ExecContext(ctx, query, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p pgHandler) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
// 	rows, err := p.db.QueryContext(ctx, query, args)
// 	if err != nil {
// 		return nil, err
// 	}

// 	row := newPgsRows(rows)

// 	return row, nil
// }

// func (p pgHandler) QueryRowContext(ctx context.Context, model, query string, args ...interface{}) repository.Row {
// 	row := p.db.QueryOneContext(ctx, model, query, params)

// 	return newPgRow(row)
// }

// type pgRow struct {
// 	row *sql.Row
// }

// func newPgRow(row *sql.Row) pgRow {
// 	return pgRow{row: row}
// }

// func (pr pgRow) Scan(dest ...interface{}) error {
// 	if err := pr.row.Scan(dest...); err != nil {
// 		return err
// 	}

// 	return nil
// }

// type pgsRows struct {
// 	rows *sql.Rows
// }

// func newPgsRows(rows *pg.Rows) pgsRows {
// 	return pgsRows{rows: rows}
// }

// func (pr pgsRows) Scan(dest ...interface{}) error {
// 	if err := pr.rows.Scan(dest...); err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (pr pgsRows) Next() bool {
// 	return pr.rows.Next()
// }

// func (pr pgsRows) Err() error {
// 	return pr.rows.Err()
// }

// func (pr pgsRows) Close() error {
// 	return pr.rows.Close()
// }

// type pgsTx struct {
// 	tx *pg.Tx
// }

// func newPgsTx(tx *pg.Tx) pgsTx {
// 	return pgsTx{tx: tx}
// }

// func (p pgsTx) ExecuteContext(ctx context.Context, query string, args ...interface{}) error {
// 	_, err := p.tx.ExecContext(ctx, query, args...)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (p pgsTx) QueryContext(ctx context.Context, query string, args ...interface{}) (repository.Rows, error) {
// 	rows, err := p.tx.QueryContext(ctx, query, args...)
// 	if err != nil {
// 		return nil, err
// 	}

// 	row := newPgsRows(rows)

// 	return row, nil
// }

// func (p pgsTx) QueryRowContext(ctx context.Context, query string, args ...interface{}) repository.Row {
// 	row := p.tx.QueryOneContext(ctx, query, args...)

// 	return newPgRow(row)
// }

// func (p pgsTx) Commit() error {
// 	return p.tx.Commit()
// }

// func (p pgsTx) Rollback() error {
// 	return p.tx.Rollback()
// }

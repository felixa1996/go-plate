package repository

import (
	"context"

	"github.com/go-pg/pg/v10"
	"gorm.io/gorm"

	_ "github.com/lib/pq"
)

type SQL interface {
	ExecuteContext(context.Context, string, ...interface{}) error
	InsertPG(context.Context, string, ...interface{}) error
	UpdatePG(context.Context, interface{}, string) error
	ExecuteContextPG(context.Context, interface{}, string, ...interface{}) error
	GetDBPG(context.Context) *pg.DB
	GetDBGorm(context.Context) *gorm.DB
	QueryContext(context.Context, string, ...interface{}) (Rows, error)
	QueryContextPG(context.Context, interface{}, string, ...interface{}) (pg.Result, error)
	QueryRowContext(context.Context, string, ...interface{}) Row
	QueryRowContextPG(context.Context, interface{}, string, ...interface{}) (pg.Result, error)
	BeginTx(ctx context.Context) (Tx, error)
}

type Rows interface {
	Scan(dest ...interface{}) error
	Next() bool
	Err() error
	Close() error
}

type Row interface {
	Scan(dest ...interface{}) error
}

type Tx interface {
	ExecuteContext(context.Context, string, ...interface{}) error
	ExecuteContextPG(context.Context, interface{}, string, ...interface{}) error
	QueryContext(context.Context, string, ...interface{}) (Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) Row
	QueryRowContextPG(context.Context, interface{}, string, ...interface{}) (pg.Result, error)
	Commit() error
	Rollback() error
}

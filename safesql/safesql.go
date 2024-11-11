package safesql

import (
	"context"
	"database/sql"

	"github.com/GA-jahida/def-prog-exercises/safesql/internal/raw"
)

type compileTimeConstant string

type TrustedSQL struct {
	s string
}

func New(text compileTimeConstant) TrustedSQL {
	return TrustedSQL{string((text))}
}

type DB struct {
	db *sql.DB
}

func (db *DB) QueryContext(ctx context.Context, query TrustedSQL, args ...any) (*sql.Rows, error) {
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

type (
	Result = sql.Result
	Rows   = sql.Rows
)

func (db *DB) ExecContext(ctx context.Context,
	query TrustedSQL, args ...any) (Result, error) {
	return db.db.ExecContext(ctx, query.s, args...)
}

func Open(driverName, dataSourceName string) (*DB, error) {
	d, err := sql.Open(driverName, dataSourceName)
	return &DB{d}, err
}

// Guaranteed to run before any package that imports safesql
func init() {
	raw.TrustedSQLCtor =
		func(unsafe string) TrustedSQL {
			return TrustedSQL{unsafe}
		}
}

func NewFromInt(i int) TrustedSQL {
	return TrustedSQL{string(i)}
}

func (t TrustedSQL) String() string {
	return t.s
}

// func TrustedSQLConcat(ss ...TrustedSQL) TrustedSQL {
// 	return TrustedSQL{string(i)}
// }

// func TrustedSQLJoin(ss []TrustedSQL, sep TrustedSQL) TrustedSQL {
// 	return TrustedSQL{string(i)}
// }

// func TrustedSQLSplit(s TrustedSQL, sep TrustedSQL) []TrustedSQL {
// 	return TrustedSQL{string(i)}
// }

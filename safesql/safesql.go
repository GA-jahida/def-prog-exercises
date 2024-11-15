package safesql

import (
	"context"
	"database/sql"
	"errors"

	auth "github.com/GA-jahida/def-prog-exercises/authentication"
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
	if auth.Must(ctx) {
		return nil, errors.New("failed")
	}
	r, err := db.db.QueryContext(ctx, query.s, args...)
	return r, err
}

type (
	Result = sql.Result
	Rows   = sql.Rows
)

func (db *DB) ExecContext(ctx context.Context,
	query TrustedSQL, args ...any) (Result, error) {
	if auth.Must(ctx) {
		return nil, errors.New("failed")
	}
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

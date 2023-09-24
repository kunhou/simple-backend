package gormhelper

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

var (
	// ErrorDuplicateValues duplicate value
	ErrorDuplicateValues = errors.New("duplicate value")
	// ErrorForeignKeyConstraint foreign key constraint
	ErrorForeignKeyConstraint = errors.New("foreign key constraint")
	// ErrorCheckConstraint check constraint
	ErrorCheckConstraint = errors.New("check constraint")
)

var (
	// https://www.postgresql.org/docs/10/errcodes-appendix.html
	codeToError = map[string]error{
		"23505": ErrorDuplicateValues,
		"23503": ErrorForeignKeyConstraint,
		"23514": ErrorCheckConstraint,
	}
)

// ParseDBError parse db error
func ParseDBError(err error) error {
	if err, ok := err.(*pgconn.PgError); ok {
		if cErr, ok := codeToError[err.Code]; ok {
			return cErr
		}
	}
	return err
}

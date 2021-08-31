package sql

import (
	"database/sql"
)

type SQLRow interface {
	Scan(dest ...interface{}) error
}

type Row struct {
	Row *sql.Row
}

func (r *Row) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

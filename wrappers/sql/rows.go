package sql

import (
	"database/sql"
)

type SQLRows interface {
	Next() bool
	NextResultSet() bool
	Err() error
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Scan(dest ...interface{}) error
	Close() error
}

type Rows struct {
	Rows *sql.Rows
}

func (rs *Rows) Next() bool {
	return rs.Rows.Next()
}

func (rs *Rows) NextResultSet() bool {
	return rs.Rows.NextResultSet()
}

func (rs *Rows) Err() error {
	return rs.Rows.Err()
}

func (rs *Rows) Columns() ([]string, error) {
	return rs.Rows.Columns()
}

func (rs *Rows) ColumnTypes() ([]*sql.ColumnType, error) {
	return rs.Rows.ColumnTypes()
}

func (rs *Rows) Scan(dest ...interface{}) error {
	return rs.Rows.Scan(dest...)
}

func (rs *Rows) Close() error {
	return rs.Rows.Close()
}

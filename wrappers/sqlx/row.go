package sqlx

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type SQLXRow interface {
	Scan(dest ...interface{}) error
	Columns() ([]string, error)
	ColumnTypes() ([]*sql.ColumnType, error)
	Err() error
	SliceScan() ([]interface{}, error)
	MapScan(dest map[string]interface{}) error
	StructScan(dest interface{}) error
}

type Row struct {
	Row *sqlx.Row
}

func (r *Row) Scan(dest ...interface{}) error {
	return r.Row.Scan(dest...)
}

func (r *Row) Columns() ([]string, error) {
	return r.Row.Columns()
}

func (r *Row) ColumnTypes() ([]*sql.ColumnType, error) {
	return r.Row.ColumnTypes()
}

func (r *Row) Err() error {
	return r.Row.Err()
}

func (r *Row) SliceScan() ([]interface{}, error) {
	return r.Row.SliceScan()
}

func (r *Row) MapScan(dest map[string]interface{}) error {
	return r.Row.MapScan(dest)
}

func (r *Row) StructScan(dest interface{}) error {
	return r.Row.StructScan(dest)
}

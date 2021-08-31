package sqlx

import (
	"database/sql"

	"github.com/jmoiron/sqlx"

	wsql "github.com/Sef1995/go-kit/wrappers/sql"
)

type SQLXRows interface {
	wsql.SQLRows
	SliceScan() ([]interface{}, error)
	MapScan(dest map[string]interface{}) error
	StructScan(dest interface{}) error
}

type Rows struct {
	Rows *sqlx.Rows
}

// sqlx.Rows

func (rs *Rows) SliceScan() ([]interface{}, error) {
	return rs.Rows.SliceScan()
}

func (rs *Rows) MapScan(dest map[string]interface{}) error {
	return rs.Rows.MapScan(dest)
}

func (rs *Rows) StructScan(dest interface{}) error {
	return rs.Rows.StructScan(dest)
}

// sql.Rows

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

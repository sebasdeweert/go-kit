package sqlx

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"github.com/jmoiron/sqlx"

	wsql "github.com/sebasdeweert/go-kit/wrappers/sql"
)

type SQLXDB interface {
	wsql.SQLDB
	DriverName() string
	MapperFunc(mf func(string) string)
	Rebind(query string) string
	Unsafe() SQLXDB
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedQuery(query string, arg interface{}) (SQLXRows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	MustBegin() SQLXTx
	Beginx() (SQLXTx, error)
	Queryx(query string, args ...interface{}) (SQLXRows, error)
	QueryRowx(query string, args ...interface{}) SQLXRow
	MustExec(query string, args ...interface{}) sql.Result
	Preparex(query string) (SQLXStmt, error)
	PrepareNamed(query string) (SQLXNamedStmt, error)
}

type DB struct {
	DB *sqlx.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sqlx.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

// sqlx.DB

func (db *DB) DriverName() string {
	return db.DB.DriverName()
}

func (db *DB) MapperFunc(mf func(string) string) {
	db.DB.MapperFunc(mf)
}

func (db *DB) Rebind(query string) string {
	return db.DB.Rebind(query)
}

func (db *DB) Unsafe() SQLXDB {
	return &DB{db.DB.Unsafe()}
}

func (db *DB) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return db.DB.BindNamed(query, arg)
}

func (db *DB) NamedQuery(query string, arg interface{}) (SQLXRows, error) {
	rows, err := db.DB.NamedQuery(query, arg)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (db *DB) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return db.DB.NamedExec(query, arg)
}

func (db *DB) Select(dest interface{}, query string, args ...interface{}) error {
	return db.DB.Select(dest, query, args...)
}

func (db *DB) Get(dest interface{}, query string, args ...interface{}) error {
	return db.DB.Get(dest, query, args...)
}

func (db *DB) MustBegin() SQLXTx {
	return &Tx{
		db.DB.MustBegin(),
	}
}

func (db *DB) Beginx() (SQLXTx, error) {
	tx, err := db.DB.Beginx()

	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

func (db *DB) Queryx(query string, args ...interface{}) (SQLXRows, error) {
	rows, err := db.DB.Queryx(query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (db *DB) QueryRowx(query string, args ...interface{}) SQLXRow {
	row := db.DB.QueryRowx(query, args...)

	if row == nil {
		return nil
	}

	return &Row{row}
}

func (db *DB) MustExec(query string, args ...interface{}) sql.Result {
	return db.DB.MustExec(query, args...)
}

func (db *DB) Preparex(query string) (SQLXStmt, error) {
	stmt, err := db.DB.Preparex(query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (db *DB) PrepareNamed(query string) (SQLXNamedStmt, error) {
	nstmt, err := db.DB.PrepareNamed(query)

	if err != nil {
		return nil, err
	}

	return &NamedStmt{nstmt}, nil
}

// sql.DB

func (db *DB) PingContext(ctx context.Context) error {
	return db.DB.PingContext(ctx)
}

func (db *DB) Ping() error {
	return db.DB.Ping()
}

func (db *DB) Close() error {
	return db.DB.Close()
}

func (db *DB) SetMaxIdleConns(n int) {
	db.DB.SetMaxIdleConns(n)
}

func (db *DB) SetMaxOpenConns(n int) {
	db.DB.SetMaxOpenConns(n)
}

func (db *DB) SetConnMaxLifetime(d time.Duration) {
	db.DB.SetConnMaxLifetime(d)
}

func (db *DB) Stats() sql.DBStats {
	return db.DB.Stats()
}

func (db *DB) PrepareContext(ctx context.Context, query string) (wsql.SQLStmt, error) {
	stmt, err := db.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &wsql.Stmt{Stmt: stmt}, nil
}

func (db *DB) Prepare(query string) (wsql.SQLStmt, error) {
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &wsql.Stmt{Stmt: stmt}, nil
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.DB.ExecContext(ctx, query, args...)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (wsql.SQLRows, error) {
	rows, err := db.DB.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (db *DB) Query(query string, args ...interface{}) (wsql.SQLRows, error) {
	rows, err := db.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) wsql.SQLRow {
	return &wsql.Row{
		Row: db.DB.QueryRowContext(ctx, query, args...),
	}
}

func (db *DB) QueryRow(query string, args ...interface{}) wsql.SQLRow {
	return &wsql.Row{
		Row: db.DB.QueryRow(query, args...),
	}
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (wsql.SQLTx, error) {
	tx, err := db.DB.BeginTx(ctx, opts)

	if err != nil {
		return nil, err
	}

	return &wsql.Tx{Tx: tx}, nil
}

func (db *DB) Begin() (wsql.SQLTx, error) {
	tx, err := db.DB.Begin()

	if err != nil {
		return nil, err
	}

	return &wsql.Tx{Tx: tx}, nil
}

func (db *DB) Driver() driver.Driver {
	return db.DB.Driver()
}

func (db *DB) Conn(ctx context.Context) (wsql.SQLConn, error) {
	conn, err := db.DB.Conn(ctx)

	if err != nil {
		return nil, err
	}

	return &wsql.Conn{Conn: conn}, nil
}

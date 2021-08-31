package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"
)

type SQLDB interface {
	PingContext(ctx context.Context) error
	Ping() error
	Close() error
	SetMaxIdleConns(n int)
	SetMaxOpenConns(n int)
	SetConnMaxLifetime(d time.Duration)
	Stats() sql.DBStats
	PrepareContext(ctx context.Context, query string) (SQLStmt, error)
	Prepare(query string) (SQLStmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error)
	Query(query string, args ...interface{}) (SQLRows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) SQLRow
	QueryRow(query string, args ...interface{}) SQLRow
	BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTx, error)
	Begin() (SQLTx, error)
	Driver() driver.Driver
	Conn(ctx context.Context) (SQLConn, error)
}

type DB struct {
	DB *sql.DB
}

func Open(driverName, dataSourceName string) (*DB, error) {
	db, err := sql.Open(driverName, dataSourceName)

	if err != nil {
		return nil, err
	}

	return &DB{db}, nil
}

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

func (db *DB) PrepareContext(ctx context.Context, query string) (SQLStmt, error) {
	stmt, err := db.DB.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (db *DB) Prepare(query string) (SQLStmt, error) {
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (db *DB) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return db.DB.ExecContext(ctx, query, args...)
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.DB.Exec(query, args...)
}

func (db *DB) QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error) {
	rows, err := db.DB.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (db *DB) Query(query string, args ...interface{}) (SQLRows, error) {
	rows, err := db.DB.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (db *DB) QueryRowContext(ctx context.Context, query string, args ...interface{}) SQLRow {
	return &Row{
		db.DB.QueryRowContext(ctx, query, args...),
	}
}

func (db *DB) QueryRow(query string, args ...interface{}) SQLRow {
	return &Row{
		db.DB.QueryRow(query, args...),
	}
}

func (db *DB) BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTx, error) {
	tx, err := db.DB.BeginTx(ctx, opts)

	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

func (db *DB) Begin() (SQLTx, error) {
	tx, err := db.DB.Begin()

	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

func (db *DB) Driver() driver.Driver {
	return db.DB.Driver()
}

func (db *DB) Conn(ctx context.Context) (SQLConn, error) {
	conn, err := db.DB.Conn(ctx)

	if err != nil {
		return nil, err
	}

	return &Conn{conn}, nil
}

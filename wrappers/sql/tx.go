package sql

import (
	"context"
	"database/sql"
)

type SQLTx interface {
	Commit() error
	Rollback() error
	PrepareContext(ctx context.Context, query string) (SQLStmt, error)
	Prepare(query string) (SQLStmt, error)
	StmtContext(ctx context.Context, stmt *sql.Stmt) SQLStmt
	Stmt(stmt *sql.Stmt) SQLStmt
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error)
	Query(query string, args ...interface{}) (SQLRows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) SQLRow
	QueryRow(query string, args ...interface{}) SQLRow
}

type Tx struct {
	Tx *sql.Tx
}

func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *Tx) PrepareContext(ctx context.Context, query string) (SQLStmt, error) {
	stmt, err := tx.Tx.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (tx *Tx) Prepare(query string) (SQLStmt, error) {
	stmt, err := tx.Tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (tx *Tx) StmtContext(ctx context.Context, stmt *sql.Stmt) SQLStmt {
	return &Stmt{
		tx.Tx.StmtContext(ctx, stmt),
	}
}

func (tx *Tx) Stmt(stmt *sql.Stmt) SQLStmt {
	return &Stmt{
		tx.Tx.Stmt(stmt),
	}
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.ExecContext(ctx, query, args...)
}

func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.Exec(query, args...)
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error) {
	rows, err := tx.Tx.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (tx *Tx) Query(query string, args ...interface{}) (SQLRows, error) {
	rows, err := tx.Tx.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) SQLRow {
	return &Row{
		tx.Tx.QueryRowContext(ctx, query, args...),
	}
}

func (tx *Tx) QueryRow(query string, args ...interface{}) SQLRow {
	return &Row{
		tx.Tx.QueryRow(query, args...),
	}
}

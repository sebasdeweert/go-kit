package sqlx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	wsql "github.com/sebasdeweert/go-kit/wrappers/sql"
)

type SQLXTx interface {
	wsql.SQLTx
	DriverName() string
	Rebind(query string) string
	Unsafe() SQLXTx
	BindNamed(query string, arg interface{}) (string, []interface{}, error)
	NamedQuery(query string, arg interface{}) (SQLXRows, error)
	NamedExec(query string, arg interface{}) (sql.Result, error)
	Select(dest interface{}, query string, args ...interface{}) error
	Queryx(query string, args ...interface{}) (SQLXRows, error)
	QueryRowx(query string, args ...interface{}) SQLXRow
	Get(dest interface{}, query string, args ...interface{}) error
	MustExec(query string, args ...interface{}) sql.Result
	Preparex(query string) (SQLXStmt, error)
	Stmtx(stmt interface{}) SQLXStmt
	NamedStmt(stmt *sqlx.NamedStmt) SQLXNamedStmt
	PrepareNamed(query string) (SQLXNamedStmt, error)
}

type Tx struct {
	Tx *sqlx.Tx
}

// sqlx.Tx

func (tx *Tx) DriverName() string {
	return tx.Tx.DriverName()
}

func (tx *Tx) Rebind(query string) string {
	return tx.Tx.Rebind(query)
}

func (tx *Tx) Unsafe() SQLXTx {
	return &Tx{
		tx.Tx.Unsafe(),
	}
}

func (tx *Tx) BindNamed(query string, arg interface{}) (string, []interface{}, error) {
	return tx.Tx.BindNamed(query, arg)
}

func (tx *Tx) NamedQuery(query string, arg interface{}) (SQLXRows, error) {
	rows, err := tx.Tx.NamedQuery(query, arg)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (tx *Tx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	return tx.Tx.NamedExec(query, arg)
}

func (tx *Tx) Select(dest interface{}, query string, args ...interface{}) error {
	return tx.Tx.Select(dest, query, args...)
}

func (tx *Tx) Queryx(query string, args ...interface{}) (SQLXRows, error) {
	rows, err := tx.Tx.Queryx(query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (tx *Tx) QueryRowx(query string, args ...interface{}) SQLXRow {
	return &Row{
		tx.Tx.QueryRowx(query, args...),
	}
}

func (tx *Tx) Get(dest interface{}, query string, args ...interface{}) error {
	return tx.Tx.Get(dest, query, args...)
}

func (tx *Tx) MustExec(query string, args ...interface{}) sql.Result {
	return tx.Tx.MustExec(query, args...)
}

func (tx *Tx) Preparex(query string) (SQLXStmt, error) {
	stmt, err := tx.Tx.Preparex(query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (tx *Tx) Stmtx(stmt interface{}) SQLXStmt {
	return &Stmt{
		tx.Tx.Stmtx(stmt),
	}
}

func (tx *Tx) NamedStmt(stmt *sqlx.NamedStmt) SQLXNamedStmt {
	return &NamedStmt{
		tx.Tx.NamedStmt(stmt),
	}
}

func (tx *Tx) PrepareNamed(query string) (SQLXNamedStmt, error) {
	nstmt, err := tx.Tx.PrepareNamed(query)

	if err != nil {
		return nil, err
	}

	return &NamedStmt{nstmt}, nil
}

// sql.Tx

func (tx *Tx) Commit() error {
	return tx.Tx.Commit()
}

func (tx *Tx) Rollback() error {
	return tx.Tx.Rollback()
}

func (tx *Tx) PrepareContext(ctx context.Context, query string) (wsql.SQLStmt, error) {
	stmt, err := tx.Tx.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &wsql.Stmt{Stmt: stmt}, nil
}

func (tx *Tx) Prepare(query string) (wsql.SQLStmt, error) {
	stmt, err := tx.Tx.Prepare(query)

	if err != nil {
		return nil, err
	}

	return &wsql.Stmt{Stmt: stmt}, nil
}

func (tx *Tx) StmtContext(ctx context.Context, stmt *sql.Stmt) wsql.SQLStmt {
	return &wsql.Stmt{
		Stmt: tx.Tx.StmtContext(ctx, stmt),
	}
}

func (tx *Tx) Stmt(stmt *sql.Stmt) wsql.SQLStmt {
	return &wsql.Stmt{
		Stmt: tx.Tx.Stmt(stmt),
	}
}

func (tx *Tx) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.ExecContext(ctx, query, args...)
}

func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.Tx.Exec(query, args...)
}

func (tx *Tx) QueryContext(ctx context.Context, query string, args ...interface{}) (wsql.SQLRows, error) {
	rows, err := tx.Tx.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (tx *Tx) Query(query string, args ...interface{}) (wsql.SQLRows, error) {
	rows, err := tx.Tx.Query(query, args...)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (tx *Tx) QueryRowContext(ctx context.Context, query string, args ...interface{}) wsql.SQLRow {
	return &wsql.Row{
		Row: tx.Tx.QueryRowContext(ctx, query, args...),
	}
}

func (tx *Tx) QueryRow(query string, args ...interface{}) wsql.SQLRow {
	return &wsql.Row{
		Row: tx.Tx.QueryRow(query, args...),
	}
}

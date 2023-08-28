package sqlx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	wsql "github.com/sebasdeweert/go-kit/wrappers/sql"
)

type SQLXStmt interface {
	wsql.SQLStmt
	Unsafe() SQLXStmt
	Select(dest interface{}, args ...interface{}) error
	Get(dest interface{}, args ...interface{}) error
	MustExec(args ...interface{}) sql.Result
	QueryRowx(args ...interface{}) SQLXRow
	Queryx(args ...interface{}) (SQLXRows, error)
}

type Stmt struct {
	Stmt *sqlx.Stmt
}

// sqlx.Stmt

func (s *Stmt) Unsafe() SQLXStmt {
	return &Stmt{
		s.Stmt.Unsafe(),
	}
}

func (s *Stmt) Select(dest interface{}, args ...interface{}) error {
	return s.Stmt.Select(dest, args...)
}

func (s *Stmt) Get(dest interface{}, args ...interface{}) error {
	return s.Stmt.Get(dest, args...)
}

func (s *Stmt) MustExec(args ...interface{}) sql.Result {
	return s.Stmt.MustExec(args...)
}

func (s *Stmt) QueryRowx(args ...interface{}) SQLXRow {
	return &Row{
		s.Stmt.QueryRowx(args...),
	}
}

func (s *Stmt) Queryx(args ...interface{}) (SQLXRows, error) {
	rows, err := s.Stmt.Queryx(args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

// sql.Stmt

func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return s.Stmt.ExecContext(ctx, args...)
}

func (s *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.Stmt.Exec(args...)
}

func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (wsql.SQLRows, error) {
	rows, err := s.Stmt.QueryContext(ctx, args...)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (s *Stmt) Query(args ...interface{}) (wsql.SQLRows, error) {
	rows, err := s.Stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) wsql.SQLRow {
	return &wsql.Row{
		Row: s.Stmt.QueryRowContext(ctx, args...),
	}
}

func (s *Stmt) QueryRow(args ...interface{}) wsql.SQLRow {
	return &wsql.Row{
		Row: s.Stmt.QueryRow(args...),
	}
}

func (s *Stmt) Close() error {
	return s.Stmt.Close()
}

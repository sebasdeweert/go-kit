package sql

import (
	"context"
	"database/sql"
)

type SQLStmt interface {
	ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error)
	Exec(args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, args ...interface{}) (SQLRows, error)
	Query(args ...interface{}) (SQLRows, error)
	QueryRowContext(ctx context.Context, args ...interface{}) SQLRow
	QueryRow(args ...interface{}) SQLRow
	Close() error
}

type Stmt struct {
	Stmt *sql.Stmt
}

func (s *Stmt) ExecContext(ctx context.Context, args ...interface{}) (sql.Result, error) {
	return s.Stmt.ExecContext(ctx, args...)
}

func (s *Stmt) Exec(args ...interface{}) (sql.Result, error) {
	return s.Stmt.Exec(args...)
}

func (s *Stmt) QueryContext(ctx context.Context, args ...interface{}) (SQLRows, error) {
	rows, err := s.Stmt.QueryContext(ctx, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (s *Stmt) Query(args ...interface{}) (SQLRows, error) {
	rows, err := s.Stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (s *Stmt) QueryRowContext(ctx context.Context, args ...interface{}) SQLRow {
	return &Row{
		s.Stmt.QueryRowContext(ctx, args...),
	}
}

func (s *Stmt) QueryRow(args ...interface{}) SQLRow {
	return &Row{
		s.Stmt.QueryRow(args...),
	}
}

func (s *Stmt) Close() error {
	return s.Stmt.Close()
}

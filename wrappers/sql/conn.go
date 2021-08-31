package sql

import (
	"context"
	"database/sql"
)

type SQLConn interface {
	PingContext(ctx context.Context) error
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) SQLRow
	PrepareContext(ctx context.Context, query string) (SQLStmt, error)
	BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTx, error)
	Close() error
}

type Conn struct {
	Conn *sql.Conn
}

func (c *Conn) PingContext(ctx context.Context) error {
	return c.Conn.PingContext(ctx)
}

func (c *Conn) ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error) {
	return c.Conn.ExecContext(ctx, query, args...)
}

func (c *Conn) QueryContext(ctx context.Context, query string, args ...interface{}) (SQLRows, error) {
	rows, err := c.Conn.QueryContext(ctx, query, args...)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (c *Conn) QueryRowContext(ctx context.Context, query string, args ...interface{}) SQLRow {
	return &Row{
		c.Conn.QueryRowContext(ctx, query, args...),
	}
}

func (c *Conn) PrepareContext(ctx context.Context, query string) (SQLStmt, error) {
	stmt, err := c.Conn.PrepareContext(ctx, query)

	if err != nil {
		return nil, err
	}

	return &Stmt{stmt}, nil
}

func (c *Conn) BeginTx(ctx context.Context, opts *sql.TxOptions) (SQLTx, error) {
	tx, err := c.Conn.BeginTx(ctx, opts)

	if err != nil {
		return nil, err
	}

	return &Tx{tx}, nil
}

func (c *Conn) Close() error {
	return c.Conn.Close()
}

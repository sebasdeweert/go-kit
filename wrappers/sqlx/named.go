package sqlx

import (
	"context"
	"database/sql"

	"github.com/jmoiron/sqlx"

	wsql "github.com/sebasdeweert/go-kit/wrappers/sql"
)

type SQLXNamedStmt interface {
	Close() error
	Exec(arg interface{}) (sql.Result, error)
	Query(arg interface{}) (wsql.SQLRows, error)
	QueryRow(arg interface{}) SQLXRow
	MustExec(arg interface{}) sql.Result
	Queryx(arg interface{}) (SQLXRows, error)
	QueryRowx(arg interface{}) SQLXRow
	Select(dest interface{}, arg interface{}) error
	Get(dest interface{}, arg interface{}) error
	Unsafe() SQLXNamedStmt
	ExecContext(ctx context.Context, arg interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, arg interface{}) (wsql.SQLRows, error)
	QueryRowContext(ctx context.Context, arg interface{}) SQLXRow
	MustExecContext(ctx context.Context, arg interface{}) sql.Result
	QueryxContext(ctx context.Context, arg interface{}) (SQLXRows, error)
	QueryRowxContext(ctx context.Context, arg interface{}) SQLXRow
	SelectContext(ctx context.Context, dest interface{}, arg interface{}) error
	GetContext(ctx context.Context, dest interface{}, arg interface{}) error
}

type NamedStmt struct {
	NamedStmt *sqlx.NamedStmt
}

func (n *NamedStmt) Close() error {
	return n.NamedStmt.Close()
}

func (n *NamedStmt) Exec(arg interface{}) (sql.Result, error) {
	return n.NamedStmt.Exec(arg)
}

func (n *NamedStmt) Query(arg interface{}) (wsql.SQLRows, error) {
	rows, err := n.NamedStmt.Query(arg)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (n *NamedStmt) QueryRow(arg interface{}) SQLXRow {
	return &Row{
		n.NamedStmt.QueryRow(arg),
	}
}

func (n *NamedStmt) MustExec(arg interface{}) sql.Result {
	return n.NamedStmt.MustExec(arg)
}

func (n *NamedStmt) Queryx(arg interface{}) (SQLXRows, error) {
	rows, err := n.NamedStmt.Queryx(arg)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (n *NamedStmt) QueryRowx(arg interface{}) SQLXRow {
	return &Row{
		n.NamedStmt.QueryRowx(arg),
	}
}

func (n *NamedStmt) Select(dest interface{}, arg interface{}) error {
	return n.NamedStmt.Select(dest, arg)
}

func (n *NamedStmt) Get(dest interface{}, arg interface{}) error {
	return n.NamedStmt.Get(dest, arg)
}

func (n *NamedStmt) Unsafe() SQLXNamedStmt {
	return &NamedStmt{
		n.NamedStmt.Unsafe(),
	}
}

func (n *NamedStmt) ExecContext(ctx context.Context, arg interface{}) (sql.Result, error) {
	return n.NamedStmt.ExecContext(ctx, arg)
}

func (n *NamedStmt) QueryContext(ctx context.Context, arg interface{}) (wsql.SQLRows, error) {
	rows, err := n.NamedStmt.QueryContext(ctx, arg)

	if err != nil {
		return nil, err
	}

	return &wsql.Rows{Rows: rows}, nil
}

func (n *NamedStmt) QueryRowContext(ctx context.Context, arg interface{}) SQLXRow {
	return &Row{
		n.NamedStmt.QueryRowContext(ctx, arg),
	}
}

func (n *NamedStmt) MustExecContext(ctx context.Context, arg interface{}) sql.Result {
	return n.NamedStmt.MustExecContext(ctx, arg)
}

func (n *NamedStmt) QueryxContext(ctx context.Context, arg interface{}) (SQLXRows, error) {
	rows, err := n.NamedStmt.QueryxContext(ctx, arg)

	if err != nil {
		return nil, err
	}

	return &Rows{rows}, nil
}

func (n *NamedStmt) QueryRowxContext(ctx context.Context, arg interface{}) SQLXRow {
	return n.QueryRowContext(ctx, arg)
}

func (n *NamedStmt) SelectContext(ctx context.Context, dest interface{}, arg interface{}) error {
	return n.NamedStmt.SelectContext(ctx, dest, arg)
}

func (n *NamedStmt) GetContext(ctx context.Context, dest interface{}, arg interface{}) error {
	return n.NamedStmt.GetContext(ctx, dest, arg)
}

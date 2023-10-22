package repo

import (
	"context"
	"database/sql"
)

const ErrMsg = "PG-error"
const TransactionCtxKey = "pg_transaction"

type conSt interface {
	ExecContext(ctx context.Context, sql string, arguments ...any) (sql.Result, error)
	QueryContext(ctx context.Context, sql string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, sql string, args ...any) *sql.Row
}

type txContainerSt struct {
	tx sql.Tx
}

func ExecContext(ctx context.Context, db *sql.DB, sql string, args ...any) (sql.Result, error) {
	conn := coalesceConn(ctx, db)
	return conn.ExecContext(ctx, sql, args)
}

func QueryContext(ctx context.Context, db *sql.DB, sql string, args ...any) (*sql.Rows, error) {
	conn := coalesceConn(ctx, db)
	return conn.QueryContext(ctx, sql, args)
}

func QueryRowContext(ctx context.Context, db *sql.DB, sql string, args ...any) *sql.Row {
	conn := coalesceConn(ctx, db)
	return conn.QueryRowContext(ctx, sql, args)
}

func getContextTransactionContainer(ctx context.Context) *txContainerSt {
	contextV := ctx.Value(TransactionCtxKey)
	if contextV == nil {
		return nil
	}

	switch tx := contextV.(type) {
	case *txContainerSt:
		return tx
	default:
		return nil
	}
}

func getContextTransaction(ctx context.Context) *sql.Tx {
	container := getContextTransactionContainer(ctx)
	if container != nil {
		return &container.tx
	}
	return nil
}

func coalesceConn(ctx context.Context, db *sql.DB) conSt {
	tx := getContextTransaction(ctx)
	if tx != nil {
		return tx
	}
	return db
}

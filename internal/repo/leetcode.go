package repo

import (
	"context"
	"database/sql"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

const ErrMsg = "PG-error"
const TransactionCtxKey = "pg_transaction"

type conSt interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}

type txContainerSt struct {
	tx sql.Tx
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

func coalesceConn(ctx context.Context, db *sql.DB) interface{} {
	tx := getContextTransaction(ctx)
	if tx != nil {
		return tx
	}
	return db
}

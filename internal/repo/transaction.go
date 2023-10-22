package repo

import (
	"context"
	"database/sql"
	"fmt"
)

type transaction struct {
	db *sql.DB
}

func newTransaction(db *sql.DB) *transaction {
	return &transaction{db: db}
}

func (u *transaction) BeginTransaction(ctx context.Context) (context.Context, error) {
	tx, err := u.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, fmt.Errorf("begin transaction: %w", err)
	}
	txContext := context.WithValue(ctx, "tx", tx)
	return txContext, nil
}

func (u *transaction) RollbackTransaction(ctx context.Context) error {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return fmt.Errorf("transaction not found in context")
	}

	err := tx.Rollback()
	if err != nil {
		return fmt.Errorf("rollback transaction: %w", err)
	}
	return nil
}

func (u *transaction) CommitTransaction(ctx context.Context) error {
	tx, ok := ctx.Value("tx").(*sql.Tx)
	if !ok {
		return fmt.Errorf("transaction not found in context")
	}

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("commit transaction: %w", err)
	}
	return nil
}

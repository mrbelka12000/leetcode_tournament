package repo

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

type Tx struct {
	db *sql.DB
}

func newTx(db *sql.DB) *Tx {
	return &Tx{
		db: db,
	}
}

type txContainerSt struct {
	tx *sql.Tx
}

func (t *Tx) ContextWithTransaction(ctx context.Context) (context.Context, error) {
	tx, err := t.db.Begin()
	if err != nil {
		return ctx, fmt.Errorf("tx begin: %w", err)
	}

	return context.WithValue(ctx, consts.TransactionCtxKey, &txContainerSt{tx: tx}), nil
}

func (t *Tx) CommitContextTransaction(ctx context.Context) error {
	tx := getContextTransaction(ctx)
	if tx == nil {
		return nil
	}

	err := tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func (t *Tx) RollbackContextTransaction(ctx context.Context) {
	tx := getContextTransaction(ctx)
	if tx == nil {
		return
	}

	_ = tx.Rollback()
}

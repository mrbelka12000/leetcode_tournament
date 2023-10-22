package transaction

import (
	"context"
)

type Transaction struct {
	transactionRepo Repo
}

func New(transactionRepo Repo) *Transaction {
	return &Transaction{
		transactionRepo: transactionRepo,
	}
}

func (e *Transaction) BeginTransaction(ctx context.Context) (context.Context, error) {
	ctx, err := e.transactionRepo.BeginTransaction(ctx)
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func (e *Transaction) RollbackTransaction(ctx context.Context) error {
	err := e.transactionRepo.RollbackTransaction(ctx)
	if err != nil {
		return err
	}
	return nil
}
func (e *Transaction) CommitTransaction(ctx context.Context) error {
	err := e.transactionRepo.CommitTransaction(ctx)
	if err != nil {
		return err
	}
	return nil
}

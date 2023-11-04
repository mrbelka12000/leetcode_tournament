package tx

import "context"

type Tx struct {
	txRepo Repo
}

func NewTx(txRepo Repo) *Tx {
	return &Tx{
		txRepo: txRepo,
	}
}

func (t *Tx) ContextWithTransaction(ctx context.Context) (context.Context, error) {
	return t.txRepo.ContextWithTransaction(ctx)
}

func (t *Tx) CommitContextTransaction(ctx context.Context) error {
	return t.txRepo.CommitContextTransaction(ctx)
}

func (t *Tx) RollbackContextTransaction(ctx context.Context) {
	t.txRepo.RollbackContextTransaction(ctx)
}

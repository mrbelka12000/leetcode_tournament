package tx

import "context"

type Repo interface {
	ContextWithTransaction(ctx context.Context) (context.Context, error)
	CommitContextTransaction(ctx context.Context) error
	RollbackContextTransaction(ctx context.Context)
}

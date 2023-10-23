package transaction

import (
	"context"
)

type (
	Repo interface {
		BeginTransaction(ctx context.Context) (context.Context, error)
		RollbackTransaction(ctx context.Context) error
		CommitTransaction(ctx context.Context) error
	}
)

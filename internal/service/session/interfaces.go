package session

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Save(ctx context.Context, obj models.Session) error
		Delete(ctx context.Context, token string) error
		Get(ctx context.Context, pars models.SessionGetPars) (models.Session, error)
		Update(ctx context.Context, obj models.Session) error
	}
)

package session

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj models.Session) error
		Delete(ctx context.Context, token string) error
		Get(ctx context.Context, token string) (models.Session, error)
	}
)

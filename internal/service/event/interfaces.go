package event

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj models.EventCU) (int64, error)
		Update(ctx context.Context, obj models.EventCU, id int64) error
		Get(ctx context.Context, pars models.EventGetPars) (models.Event, error)
		List(ctx context.Context, pars models.EventListPars) ([]models.Event, int64, error)
	}
)

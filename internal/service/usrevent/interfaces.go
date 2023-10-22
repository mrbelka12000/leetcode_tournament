package usrevent

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj models.UsrEventCU) (int64, error)
		Update(ctx context.Context, obj models.UsrEventCU, id int64) error
		Get(ctx context.Context, pars models.UsrEventGetPars) (models.UsrEvent, error)
		List(ctx context.Context, pars models.UsrEventListPars) ([]models.UsrEvent, int64, error)
		GetUsrEvents(ctx context.Context, pars models.UsrGetEventsPars) ([]models.Event, int64, error)
	}
)

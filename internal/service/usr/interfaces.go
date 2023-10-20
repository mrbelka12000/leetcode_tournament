package usr

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj models.UsrCU) (int64, error)
		Update(ctx context.Context, obj models.UsrCU, id int64) error
		Get(ctx context.Context, pars models.UsrGetPars) (models.Usr, error)
		List(ctx context.Context, pars models.UsrListPars) ([]models.Usr, int64, error)
	}
)

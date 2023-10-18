package score

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj *models.ScoreCU) (int64, error)
		Update(ctx context.Context, obj *models.ScoreCU, id int64) error
		Get(ctx context.Context, pars *models.ScoreGetPars) (*models.Score, error)
		List(ctx context.Context, pars *models.ScoreListPars) ([]*models.Score, int64, error)
	}
)

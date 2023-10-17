package tournament

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj *models.TournamentCU) (int64, error)
		Update(ctx context.Context, obj *models.TournamentCU, id int64) error
		Get(ctx context.Context, pars *models.TournamentGetPars) (*models.Tournament, error)
		List(ctx context.Context, pars *models.TournamentListPars) ([]*models.Tournament, int64, error)
	}
)

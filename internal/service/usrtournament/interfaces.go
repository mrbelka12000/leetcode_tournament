package usrtournament

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Repo interface {
		Create(ctx context.Context, obj *models.UsrTournamentCU) (int64, error)
		Update(ctx context.Context, obj *models.UsrTournamentCU, id int64) error
		Get(ctx context.Context, pars *models.UsrTournamentGetPars) (*models.UsrTournament, error)
		List(ctx context.Context, pars *models.UsrTournamentListPars) ([]*models.UsrTournament, int64, error)
	}
)

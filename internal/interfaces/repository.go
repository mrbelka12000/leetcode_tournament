package interfaces

import (
	"context"

	models2 "github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type (
	Tournament interface {
		Create(ctx context.Context, obj *models2.TournamentCU) (int64, error)
		Update(ctx context.Context, obj *models2.TournamentCU, id int64) error
		Get(ctx context.Context, pars *models2.TournamentGetPars) (*models2.Tournament, error)
		List(ctx context.Context, pars *models2.TournamentListPars) ([]*models2.Tournament, int64, error)
	}
	UsrEvent interface {
		Create(ctx context.Context, obj *models2.UsrEventCU) (int64, error)
		Update(ctx context.Context, obj *models2.UsrEventCU, id int64) error
		Get(ctx context.Context, pars *models2.UsrEventGetPars) (*models2.UsrEvent, error)
		List(ctx context.Context, pars *models2.UsrEventListPars) ([]*models2.UsrEvent, int64, error)
		GetUsrEvents(ctx context.Context, pars *models2.UsrGetEventsPars) ([]*models2.Event, int64, error)
	}
	UsrTournament interface {
		Create(ctx context.Context, obj *models2.UsrTournamentCU) (int64, error)
		Update(ctx context.Context, obj *models2.UsrTournamentCU, id int64) error
		Get(ctx context.Context, pars *models2.UsrTournamentGetPars) (*models2.UsrTournament, error)
		List(ctx context.Context, pars *models2.UsrTournamentListPars) ([]*models2.UsrTournament, int64, error)
	}
)

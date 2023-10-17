package interfaces

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type (
	UsrRepo interface {
		Create(ctx context.Context, obj *models.UsrCU) (int64, error)
		Update(ctx context.Context, obj *models.UsrCU, id int64) error
		Get(ctx context.Context, pars *models.UsrGetPars) (*models.Usr, error)
		List(ctx context.Context, pars *models.UsrListPars) ([]*models.Usr, int64, error)
	}

	Score interface {
		Create(ctx context.Context, obj *models.ScoreCU) (int64, error)
		Update(ctx context.Context, obj *models.ScoreCU, id int64) error
		Get(ctx context.Context, pars *models.ScoreGetPars) (*models.Score, error)
		List(ctx context.Context, pars *models.ScoreListPars) ([]*models.Score, int64, error)
	}

	Event interface {
		Create(ctx context.Context, obj *models.EventCU) (int64, error)
		Update(ctx context.Context, obj *models.EventCU, id int64) error
		Get(ctx context.Context, pars *models.EventGetPars) (*models.Event, error)
		List(ctx context.Context, pars *models.EventListPars) ([]*models.Event, int64, error)
	}
	Tournament interface {
		Create(ctx context.Context, obj *models.TournamentCU) (int64, error)
		Update(ctx context.Context, obj *models.TournamentCU, id int64) error
		Get(ctx context.Context, pars *models.TournamentGetPars) (*models.Tournament, error)
		List(ctx context.Context, pars *models.TournamentListPars) ([]*models.Tournament, int64, error)
	}
	UsrEvent interface {
		Create(ctx context.Context, obj *models.UsrEventCU) (int64, error)
		Update(ctx context.Context, obj *models.UsrEventCU, id int64) error
		Get(ctx context.Context, pars *models.UsrEventGetPars) (*models.UsrEvent, error)
		List(ctx context.Context, pars *models.UsrEventListPars) ([]*models.UsrEvent, int64, error)
	}
	UsrTournament interface {
		Create(ctx context.Context, obj *models.UsrTournamentCU) (int64, error)
		Update(ctx context.Context, obj *models.UsrTournamentCU, id int64) error
		Get(ctx context.Context, pars *models.UsrTournamentGetPars) (*models.UsrTournament, error)
		List(ctx context.Context, pars *models.UsrTournamentListPars) ([]*models.UsrTournament, int64, error)
	}
)

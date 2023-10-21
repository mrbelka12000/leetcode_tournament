package tournament

import (
	"context"
	"fmt"
	"time"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/ptr"
)

type Tournament struct {
	tournamentRepo Repo
}

func New(tournamentRepo Repo) *Tournament {
	return &Tournament{
		tournamentRepo: tournamentRepo,
	}
}

// Build ..
func (e *Tournament) Build(ctx context.Context, obj models.TournamentCU) (int64, error) {
	obj.StatusID = ptr.TournamentStatusPointer(consts.TournamentStatusCreated)
	obj.PrizePool = pointer.ToInt(0)

	err := e.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}

	return e.tournamentRepo.Create(ctx, obj)
}

// Update ..
func (e *Tournament) Update(ctx context.Context, obj models.TournamentCU, id int64) error {
	err := e.validateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	return e.tournamentRepo.Update(ctx, obj, id)
}

// Get ..
func (e *Tournament) Get(ctx context.Context, pars models.TournamentGetPars, errNE bool) (models.Tournament, error) {
	tournament, err := e.tournamentRepo.Get(ctx, pars)
	if err != nil {
		return models.Tournament{}, fmt.Errorf("tournament get from db: %w", err)
	}
	if errNE && tournament.ID == 0 {
		return models.Tournament{}, errs.ErrTournamentNotFound
	}

	return tournament, nil
}

// List ..
func (e *Tournament) List(ctx context.Context, pars models.TournamentListPars) ([]models.Tournament, int64, error) {
	return e.tournamentRepo.List(ctx, pars)
}

// validateCU ..
func (e *Tournament) validateCU(ctx context.Context, obj models.TournamentCU, id int64) error {
	forCreate := id == 0

	if forCreate && obj.UsrID == nil {
		return errs.ErrUsrIDNotFound
	}
	if forCreate && obj.StartTime == nil {
		return errs.ErrStartTimeNotFound
	}
	if obj.StartTime != nil {
		if time.Now().After(*obj.StartTime) {
			return errs.ErrInvalidStartTime
		}
	}

	if forCreate && obj.EndTime == nil {
		return errs.ErrEndTimeNotFound
	}
	if obj.EndTime != nil {
		if time.Now().After(*obj.EndTime) {
			return errs.ErrInvalidEndTime
		}

		if obj.StartTime.After(*obj.EndTime) {
			return errs.ErrInvalidEndTime
		}
	}

	if forCreate && obj.Goal == nil {
		return errs.ErrGoalNotFound
	}
	if obj.Goal != nil {
		if *obj.Goal <= 0 {
			return errs.ErrInvalidGoal
		}
	}

	if forCreate && obj.StatusID == nil {
		return errs.ErrTournamentStatusNotFound
	}
	if obj.StatusID != nil {
		if !consts.IsValidTournamentStatus(*obj.StatusID) {
			return errs.ErrInvalidTournamentStatus
		}
	}
	if forCreate && obj.Cost == nil {
		return errs.ErrTournamentCostNotFound
	}
	if obj.Cost != nil {
		if *obj.Cost <= 0 {
			return errs.ErrInvalidTournamentCost
		}
	}

	return nil
}

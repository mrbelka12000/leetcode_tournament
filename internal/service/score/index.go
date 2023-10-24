package score

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/client/leetcode"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type Score struct {
	scoreRepo     Repo
	leetCodeStats LeetCodeStats
}

func New(scoreRepo Repo, ls LeetCodeStats) *Score {
	return &Score{
		scoreRepo:     scoreRepo,
		leetCodeStats: ls,
	}
}

// Stats ..
func (u *Score) Stats(ctx context.Context, nickname string) (resp leetcode.LCGetProblemsSolvedResp, err error) {
	resp, err = u.leetCodeStats.Stats(ctx, nickname)
	if err != nil {
		return leetcode.LCGetProblemsSolvedResp{}, fmt.Errorf("get stats from leetcode: %w", err)
	}

	return resp, nil
}

// Build ..
func (u *Score) Build(ctx context.Context, obj models.ScoreCU) (int64, error) {
	obj.Active = pointer.ToBool(false)

	err := u.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}

	id, err := u.scoreRepo.Create(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("create score: %w", err)
	}

	return id, nil
}

// Update ..
func (u *Score) Update(ctx context.Context, obj models.ScoreCU, id int64) error {
	err := u.validateCU(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("validate: %w", err)
	}

	err = u.scoreRepo.Update(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("update score: %w", err)
	}

	return nil
}

// Get ..
func (u *Score) Get(ctx context.Context, pars models.ScoreGetPars, errNE bool) (models.Score, error) {
	score, err := u.scoreRepo.Get(ctx, pars)
	if err != nil {
		return models.Score{}, err
	}
	if errNE && score.ID == 0 {
		return models.Score{}, errs.ErrNotFound
	}

	return score, err
}

// List ..
func (u *Score) List(ctx context.Context, pars models.ScoreListPars) ([]models.Score, int64, error) {
	resp, id, err := u.scoreRepo.List(ctx, pars)
	if err != nil {
		return []models.Score{}, 0, err
	}

	return resp, id, nil
}

// validateCU ..
func (u *Score) validateCU(ctx context.Context, obj models.ScoreCU, id int64) error {
	forCreate := id == 0

	if forCreate && obj.UsrID == nil {
		return errs.ErrUsrIDNotFound
	}

	if forCreate && obj.Active == nil {
		return errs.ErrActiveNotFound
	}

	return nil
}

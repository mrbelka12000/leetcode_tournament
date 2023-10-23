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

func (u *Score) Stats(ctx context.Context, nickname string) (resp leetcode.LCGetProblemsSolvedResp, err error) {
	resp, err = u.leetCodeStats.Stats(ctx, nickname)
	if err != nil {
		return leetcode.LCGetProblemsSolvedResp{}, err
	}
	return resp, nil
}

func (u *Score) Build(ctx context.Context, obj *models.ScoreCU) (int64, error) {
	obj.Active = pointer.ToBool(true)

	err := u.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}
	id, err := u.scoreRepo.Create(ctx, obj)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (u *Score) Update(ctx context.Context, obj *models.ScoreCU, id int64) error {
	err := u.validateCU(ctx, obj, id)
	if err != nil {
		return err
	}
	err = u.scoreRepo.Update(ctx, obj, id)
	if err != nil {
		return err
	}
	return nil

}
func (u *Score) Get(ctx context.Context, pars *models.ScoreGetPars) (*models.Score, error) {
	resp, err := u.scoreRepo.Get(ctx, pars)
	if err != nil {
		return nil, err
	}
	return resp, err

}
func (u *Score) List(ctx context.Context, pars *models.ScoreListPars) ([]*models.Score, int64, error) {
	resp, id, err := u.scoreRepo.List(ctx, pars)
	if err != nil {
		return []*models.Score{}, 0, err
	}
	return resp, id, nil

}

func (u *Score) validateCU(ctx context.Context, obj *models.ScoreCU, id int64) error {
	forCreate := id == 0

	if forCreate && obj.UsrID == nil {
		return errs.ErrUsrIDNotFound
	}

	if obj.UsrID != nil {
		_, err := u.Get(ctx, &models.ScoreGetPars{
			ID:    &id,
			UsrID: obj.UsrID,
		})
		if err != nil {
			return errs.ErrPermissionDenied
		}
	}

	if forCreate && obj.Active == nil {
		return errs.ErrActiveNotFound
	}

	return nil
}

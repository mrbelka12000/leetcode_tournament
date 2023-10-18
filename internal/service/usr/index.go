package usr

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type Usr struct {
	usrRepo       Repo
	leetCodeStats LeetCodeStats
}

func New(usrRepo Repo, leetCodeStats LeetCodeStats) *Usr {
	return &Usr{
		usrRepo:       usrRepo,
		leetCodeStats: leetCodeStats,
	}
}

func (u *Usr) Build(ctx context.Context, obj *models.UsrCU) (int64, error) {
	err := u.validateCU(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("validate CU: %w", err)
	}

	return u.usrRepo.Create(ctx, obj)
}

func (u *Usr) Update(ctx context.Context, obj *models.UsrCU, id int64) error {
	err := u.validateCU(ctx, obj)
	if err != nil {
		return err
	}

	return u.usrRepo.Update(ctx, obj, id)
}

func (u *Usr) Get(ctx context.Context, pars *models.UsrGetPars) (*models.Usr, error) {
	return u.usrRepo.Get(ctx, pars)
}

func (u *Usr) List(ctx context.Context, pars *models.UsrListPars) ([]*models.Usr, int64, error) {
	return u.usrRepo.List(ctx, pars)
}

func (u *Usr) validateCU(ctx context.Context, obj *models.UsrCU) error {

	return nil
}

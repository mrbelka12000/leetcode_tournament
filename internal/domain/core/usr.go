package core

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/interfaces"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type usr struct {
	usrRepo       interfaces.UsrRepo
	leetCodeStats interfaces.LeetCodeStats
}

func newUsr(usrRepo interfaces.UsrRepo, leetCodeStats interfaces.LeetCodeStats) *usr {
	return &usr{
		usrRepo:       usrRepo,
		leetCodeStats: leetCodeStats,
	}
}

func (u *usr) Build(ctx context.Context, obj *models.UsrCU) (int64, error) {
	err := u.validateCU(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("validate CU: %w", err)
	}

	return u.usrRepo.Create(ctx, obj)
}

func (u *usr) Update(ctx context.Context, obj *models.UsrCU, id int64) error {
	err := u.validateCU(ctx, obj)
	if err != nil {
		return err
	}

	return u.usrRepo.Update(ctx, obj, id)
}

func (u *usr) Get(ctx context.Context, pars *models.UsrGetPars) (*models.Usr, error) {
	return u.usrRepo.Get(ctx, pars)
}

func (u *usr) List(ctx context.Context, pars *models.UsrListPars) ([]*models.Usr, int64, error) {
	return u.usrRepo.List(ctx, pars)
}

func (u *usr) validateCU(ctx context.Context, obj *models.UsrCU) error {

	return nil
}

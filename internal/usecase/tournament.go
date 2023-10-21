package usecase

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

// TournamentCreate ..
func (uc *UseCase) TournamentCreate(ctx context.Context, obj models.TournamentCU) (int64, error) {
	token := ctx.Value(consts.CKey).(string)
	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return 0, fmt.Errorf("get session by token: %w", err)
	}

	if ses.TypeID != consts.UsrTypeAdmin && ses.TypeID != consts.UsrTypeClient {
		return 0, errs.ErrPermissionDenied
	}

	obj.UsrID = &ses.UsrID

	return uc.cr.Tournament.Build(ctx, obj)
}

// TournamentUpdate ..
func (uc *UseCase) TournamentUpdate(ctx context.Context, obj models.TournamentCU, id int64) error {
	token := ctx.Value(consts.CKey).(string)
	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return fmt.Errorf("get session by token: %w", err)
	}
	if ses.TypeID != consts.UsrTypeAdmin {
		return errs.ErrPermissionDenied
	}

	err = uc.cr.Tournament.Update(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("event update: %w", err)
	}

	return nil
}

// TournamentGet ..
func (uc *UseCase) TournamentGet(ctx context.Context, pars models.TournamentGetPars) (models.Tournament, error) {
	return uc.cr.Tournament.Get(ctx, pars, true)
}

// TournamentList ..
func (uc *UseCase) TournamentList(ctx context.Context, pars models.TournamentListPars) ([]models.Tournament, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.Tournament.List(ctx, pars)
}

package usecase

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

// UsrEventCreate ..
func (uc *UseCase) UsrEventCreate(ctx context.Context, obj models.UsrEventCU) (int64, error) {
	ctx, err := uc.cr.Tx.ContextWithTransaction(ctx)
	if err != nil {
		return 0, err
	}
	defer uc.cr.Tx.RollbackContextTransaction(ctx)

	token := ctx.Value(consts.CKey).(string)
	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return 0, fmt.Errorf("get session by token: %w", err)
	}
	obj.UsrID = &ses.UsrID

	usr, err := uc.cr.Usr.Get(ctx, models.UsrGetPars{
		ID: &ses.UsrID,
	}, true)
	if err != nil {
		return 0, fmt.Errorf("get usr: %w", err)
	}

	score, err := uc.cr.Score.Get(ctx, models.ScoreGetPars{
		UsrID: &ses.UsrID,
	}, true)
	if err != nil {
		return 0, fmt.Errorf("get score: %w", err)
	}

	stats, err := uc.cr.Score.Stats(ctx, usr.Username)
	if err != nil {
		return 0, fmt.Errorf("get stats: %w", err)
	}

	point := models.Points(stats)
	err = uc.cr.Score.Update(ctx, models.ScoreCU{
		Footprint: &point,
		Active:    pointer.ToBool(true),
	}, score.ID)
	if err != nil {
		return 0, fmt.Errorf("score update: %w", err)
	}

	id, err := uc.cr.UsrEvent.Build(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("usr event build: %w", err)
	}

	err = uc.cr.Tx.CommitContextTransaction(ctx)
	if err != nil {
		return 0, err
	}

	return id, nil
}

// UsrEventUpdate ..
func (uc *UseCase) UsrEventUpdate(ctx context.Context, obj models.UsrEventCU, id int64) error {
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

	err = uc.cr.UsrEvent.Update(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("usr event update: %w", err)
	}

	return nil
}

// UsrEventGet ..
func (uc *UseCase) UsrEventGet(ctx context.Context, pars models.UsrEventGetPars) (models.UsrEvent, error) {
	return uc.cr.UsrEvent.Get(ctx, pars, true)
}

// UsrEventList ..
func (uc *UseCase) UsrEventList(ctx context.Context, pars models.UsrEventListPars) ([]models.UsrEvent, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.UsrEvent.List(ctx, pars)
}

package usecase

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

// EventCreate ..
func (uc *UseCase) EventCreate(ctx context.Context, obj models.EventCU) (int64, error) {
	token := ctx.Value(consts.CKey).(string)
	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return 0, fmt.Errorf("get session by token: %w", err)
	}
	obj.UsrID = &ses.UsrID

	return uc.cr.Event.Build(ctx, obj)
}

// EventUpdate ..
func (uc *UseCase) EventUpdate(ctx context.Context, obj models.EventCU, id int64) error {
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

	err = uc.cr.Event.Update(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("event update: %w", err)
	}

	return nil
}

// EventGet ..
func (uc *UseCase) EventGet(ctx context.Context, pars models.EventGetPars) (models.Event, error) {
	return uc.cr.Event.Get(ctx, pars, true)
}

// EventList ..
func (uc *UseCase) EventList(ctx context.Context, pars models.EventListPars) ([]models.Event, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.Event.List(ctx, pars)
}

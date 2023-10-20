package usecase

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
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

	// validate is event this usr`s event
	{
		_, err := uc.cr.Event.Get(ctx, models.EventGetPars{
			ID:    &id,
			UsrID: &ses.UsrID,
		}, true)
		if err != nil {
			return errs.ErrPermissionDenied
		}
	}

	err = uc.cr.Event.Update(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("event update: %w", err)
	}

	return nil
}

func (uc *UseCase) EventGet(ctx context.Context, pars models.EventGetPars) (models.Event, error) {
	return uc.cr.Event.Get(ctx, pars, true)
}

func (uc *UseCase) EventList(ctx context.Context, pars models.EventListPars) ([]models.Event, int64, error) {
	return uc.cr.Event.List(ctx, pars)
}

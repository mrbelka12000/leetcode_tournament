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
func (uc *UseCase) EventGet(ctx context.Context, pars models.EventGetPars) (models.EventPage, error) {
	var (
		eventPage models.EventPage
		err       error
	)

	eventPage.Event, err = uc.cr.Event.Get(ctx, pars, true)
	if err != nil {
		return models.EventPage{}, fmt.Errorf("get event: %w", err)
	}

	usrEvents, _, err := uc.cr.UsrEvent.List(ctx, models.UsrEventListPars{
		UsrEventGetPars: models.UsrEventGetPars{
			EventID: pars.ID,
		},
	})
	if err != nil {
		return models.EventPage{}, fmt.Errorf("get usr events: %w", err)
	}

	eventPage.Usrs = make([]models.Usr, len(usrEvents))
	for i := 0; i < len(usrEvents); i++ {
		eventPage.Usrs[i], err = uc.cr.Usr.Get(ctx, models.UsrGetPars{
			ID:       &usrEvents[i].UsrID,
			HideInfo: true,
		}, true)
		if err != nil {
			return models.EventPage{}, fmt.Errorf("get usr: %w", err)
		}

		eventPage.Usrs[i].Score, err = uc.cr.Score.Get(ctx, models.ScoreGetPars{
			UsrID: &usrEvents[i].UsrID,
		}, true)
		if err != nil {
			return models.EventPage{}, fmt.Errorf("get score: %w", err)
		}
	}

	return eventPage, nil
}

// EventList ..
func (uc *UseCase) EventList(ctx context.Context, pars models.EventListPars) ([]models.Event, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.Event.List(ctx, pars)
}

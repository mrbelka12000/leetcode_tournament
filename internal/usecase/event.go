package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/ptr"
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

// EventPageGet ..
func (uc *UseCase) EventPageGet(ctx context.Context, pars models.EventGetPars) (models.EventPage, error) {
	var (
		eventPage models.EventPage
		err       error
		usrID     int64
	)

	token, ok := ctx.Value(consts.CKey).(string)
	if ok {
		ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
			Token: &token,
		})
		if err != nil {
			return eventPage, fmt.Errorf("get session by token: %w", err)
		}

		usrID = ses.UsrID
	}

	eventPage.Event, err = uc.EventGet(ctx, pars)
	if err != nil {
		return models.EventPage{}, fmt.Errorf("get event: %w", err)
	}

	usrEvents, _, err := uc.cr.UsrEvent.List(ctx, models.UsrEventListPars{
		UsrEventGetPars: models.UsrEventGetPars{
			EventID: pars.ID,
		},
	})
	if err != nil {
		return eventPage, fmt.Errorf("get usr events: %w", err)
	}

	eventPage.Usrs = make([]models.Usr, len(usrEvents))
	for i := 0; i < len(usrEvents); i++ {
		eventPage.Usrs[i], err = uc.cr.Usr.Get(ctx, models.UsrGetPars{
			ID:       &usrEvents[i].UsrID,
			HideInfo: true,
		}, true)
		if err != nil {
			return eventPage, fmt.Errorf("get usr: %w", err)
		}

		eventPage.Usrs[i].Score, err = uc.cr.Score.Get(ctx, models.ScoreGetPars{
			UsrID: &usrEvents[i].UsrID,
		}, true)
		if err != nil {
			return eventPage, fmt.Errorf("get score: %w", err)
		}

		if usrEvents[i].UsrID == usrID {
			eventPage.IsParticipating = true
		}
	}

	return eventPage, nil
}

func (uc *UseCase) EventGet(ctx context.Context, pars models.EventGetPars) (models.Event, error) {
	event, err := uc.cr.Event.Get(ctx, pars, true)
	if err != nil {
		return models.Event{}, err
	}

	event.Usr, err = uc.cr.Usr.Get(ctx, models.UsrGetPars{
		ID:       &event.UsrID,
		HideInfo: true,
	}, true)
	if err != nil {
		return models.Event{}, fmt.Errorf("get usr: %w", err)
	}

	return event, nil
}

// EventList ..
func (uc *UseCase) EventList(ctx context.Context, pars models.EventListPars) ([]models.Event, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.Event.List(ctx, pars)
}

func (uc *UseCase) EventFinish(ctx context.Context) error {
	ctx, err := uc.cr.Tx.ContextWithTransaction(ctx)
	if err != nil {
		return err
	}
	defer uc.cr.Tx.RollbackContextTransaction(ctx)

	events, _, err := uc.cr.Event.List(ctx, models.EventListPars{
		StatusID: ptr.EventStatusPointer(consts.EventStatusStarted),
	})
	if err != nil {
		return fmt.Errorf("get events: %w", err)
	}

	for i := 0; i < len(events); i++ {
		if events[i].EndTime.After(time.Now()) {
			continue
		}

		usrEvents, _, err := uc.cr.UsrEvent.List(ctx, models.UsrEventListPars{
			UsrEventGetPars: models.UsrEventGetPars{
				EventID: &events[i].ID,
			},
		})
		if err != nil {
			return fmt.Errorf("get usr events: %w", err)
		}

		var (
			maxTotal   uint64
			usrEventID int64
		)

	USREVENT:
		for j := 0; j < len(usrEvents); j++ {
			usr, err := uc.UsrGet(ctx, models.UsrGetPars{
				ID:        &usrEvents[j].UsrID,
				WithScore: true,
			})
			if err != nil {
				return fmt.Errorf("get usr with score: %w", err)
			}

			solved := usr.Score.Current.Total - usr.Score.Footprint.Total
			if solved < events[i].Goal {
				continue
			}

			switch events[i].Condition {
			case consts.EventConditionOnMax:
				if solved > maxTotal {
					maxTotal = solved
					usrEventID = usrEvents[i].ID
				}
			case consts.EventConditionOnFirst:
				err = uc.cr.UsrEvent.Update(ctx, models.UsrEventCU{
					Winner: pointer.ToBool(true),
				}, usrEvents[i].ID)
				if err != nil {
					return fmt.Errorf("winner set to usr event: %w", err)
				}

				break USREVENT
			case consts.EventConditionOnTimeExceed:
				err = uc.cr.UsrEvent.Update(ctx, models.UsrEventCU{
					Winner: pointer.ToBool(true),
				}, usrEvents[i].ID)
				if err != nil {
					return fmt.Errorf("winner set to usr event: %w", err)
				}
			}
		}

		if usrEventID != 0 {
			err = uc.cr.UsrEvent.Update(ctx, models.UsrEventCU{
				Winner: pointer.ToBool(true),
			}, usrEventID)
			if err != nil {
				return fmt.Errorf("winner set to usr event: %w", err)
			}
		}

		err = uc.cr.Event.Update(ctx, models.EventCU{
			StatusID: ptr.EventStatusPointer(consts.EventStatusFinished),
		}, events[i].ID)
		if err != nil {
			return fmt.Errorf("event finish: %w", err)
		}

	}
	err = uc.cr.Tx.CommitContextTransaction(ctx)
	if err != nil {
		return err
	}

	uc.log.Info().Msg("events finisher successfully worked")

	return nil
}

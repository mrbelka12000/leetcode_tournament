package event

import (
	"context"
	"fmt"
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/ptr"
)

type Event struct {
	eventRepo Repo
}

func New(eventRepo Repo) *Event {
	return &Event{
		eventRepo: eventRepo,
	}
}

// Build ..
func (e *Event) Build(ctx context.Context, obj models.EventCU) (int64, error) {
	obj.StatusID = ptr.EventStatusPointer(consts.EventStatusCreated)

	err := e.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}

	return e.eventRepo.Create(ctx, obj)
}

// Update ..
func (e *Event) Update(ctx context.Context, obj models.EventCU, id int64) error {
	err := e.validateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	return e.eventRepo.Update(ctx, obj, id)
}

// Get ..
func (e *Event) Get(ctx context.Context, pars models.EventGetPars, errNE bool) (models.Event, error) {
	event, err := e.eventRepo.Get(ctx, pars)
	if err != nil {
		return models.Event{}, fmt.Errorf("event get from db: %w", err)
	}
	if errNE && event.ID == 0 {
		return models.Event{}, errs.ErrEventNotFound
	}

	return event, nil
}

// List ..
func (e *Event) List(ctx context.Context, pars models.EventListPars) ([]models.Event, int64, error) {
	return e.eventRepo.List(ctx, pars)
}

// validateCU ..
func (e *Event) validateCU(ctx context.Context, obj models.EventCU, id int64) error {
	forCreate := id == 0

	if forCreate && obj.UsrID == nil {
		return errs.ErrUsrIDNotFound
	}
	if forCreate && obj.StartTime == nil {
		return errs.ErrStartTimeNotFound
	}
	if obj.StartTime != nil {
		if time.Now().After(*obj.StartTime) {
			return errs.ErrInvalidStartTime
		}
	}

	if forCreate && obj.EndTime == nil {
		return errs.ErrEndTimeNotFound
	}
	if obj.EndTime != nil {
		if time.Now().After(*obj.EndTime) {
			return errs.ErrInvalidEndTime
		}

		if obj.StartTime.After(*obj.EndTime) {
			return errs.ErrInvalidEndTime
		}
	}

	if forCreate && obj.Goal == nil {
		return errs.ErrGoalNotFound
	}
	if obj.Goal != nil {
		if *obj.Goal <= 0 {
			return errs.ErrInvalidGoal
		}
	}

	if forCreate && obj.Condition == nil {
		return errs.ErrEventConditionNotFound
	}
	if obj.Condition != nil {
		if !consts.IsValidEventCondition(*obj.Condition) {
			return errs.ErrInvalidEventCondition
		}
	}

	if forCreate && obj.StatusID == nil {
		return errs.ErrEventStatusNotFound
	}
	if obj.StatusID != nil {
		if !consts.IsValidEventStatus(*obj.StatusID) {
			return errs.ErrInvalidEventStatus
		}
	}
	return nil
}

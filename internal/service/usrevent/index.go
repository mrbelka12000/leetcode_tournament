package usrevent

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type UsrEvent struct {
	usrEventRepo Repo
}

func New(usrEventRepo Repo) *UsrEvent {
	return &UsrEvent{
		usrEventRepo: usrEventRepo,
	}
}

func (u *UsrEvent) Build(ctx context.Context, obj models.UsrEventCU) (int64, error) {
	obj.Active = pointer.ToBool(true)
	obj.Winner = pointer.ToBool(false)

	err := u.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}

	id, err := u.usrEventRepo.Create(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("create usr event: %w", err)
	}

	return id, nil
}

func (u *UsrEvent) Update(ctx context.Context, obj models.UsrEventCU, id int64) error {
	err := u.validateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	return u.usrEventRepo.Update(ctx, obj, id)
}

func (u *UsrEvent) Get(ctx context.Context, pars models.UsrEventGetPars, errNE bool) (models.UsrEvent, error) {
	usr, err := u.usrEventRepo.Get(ctx, pars)
	if err != nil {
		return models.UsrEvent{}, fmt.Errorf("usr event get from db: %w", err)
	}
	if errNE && usr.ID == 0 {
		return models.UsrEvent{}, errs.ErrNotFound
	}

	return usr, nil
}

func (u *UsrEvent) List(ctx context.Context, pars models.UsrEventListPars) ([]models.UsrEvent, int64, error) {
	return u.usrEventRepo.List(ctx, pars)
}

func (u *UsrEvent) GetUsrEvents(ctx context.Context, pars models.UsrGetEventsPars) ([]models.Event, int64, error) {
	return u.usrEventRepo.GetUsrEvents(ctx, pars)
}

func (u *UsrEvent) validateCU(ctx context.Context, obj models.UsrEventCU, id int64) error {
	forCreate := id == 0

	if forCreate && obj.UsrID == nil {
		return errs.ErrUsrIDNotFound
	}
	if forCreate && obj.EventID == nil {
		return errs.ErrEventIDNotFound
	}

	if forCreate && obj.UsrID != nil && obj.EventID != nil {
		_, err := u.Get(ctx, models.UsrEventGetPars{
			UsrID:   obj.UsrID,
			EventID: obj.EventID,
		}, false)
		if err != nil {
			return errs.ErrPermissionDenied
		}
	}

	if forCreate && obj.Active == nil {
		return errs.ErrActiveNotFound
	}

	return nil
}

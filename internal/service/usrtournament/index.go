package usrtournament

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"

	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

type UsrTournament struct {
	usrTourRepo Repo
}

func New(usrTourRepo Repo) *UsrTournament {
	return &UsrTournament{
		usrTourRepo: usrTourRepo,
	}
}

func (u *UsrTournament) Build(ctx context.Context, obj models.UsrTournamentCU) (int64, error) {
	obj.Active = pointer.ToBool(true)
	obj.Winner = pointer.ToBool(false)

	err := u.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, fmt.Errorf("validate: %w", err)
	}

	id, err := u.usrTourRepo.Create(ctx, obj)
	if err != nil {
		return 0, fmt.Errorf("create usr event: %w", err)
	}

	return id, nil
}

func (u *UsrTournament) Update(ctx context.Context, obj models.UsrTournamentCU, id int64) error {
	err := u.validateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	return u.usrTourRepo.Update(ctx, obj, id)
}

func (u *UsrTournament) Get(ctx context.Context, pars models.UsrTournamentGetPars, errNE bool) (models.UsrTournament, error) {
	usr, err := u.usrTourRepo.Get(ctx, pars)
	if err != nil {
		return models.UsrTournament{}, fmt.Errorf("usr event get from db: %w", err)
	}
	if errNE && usr.ID == 0 {
		return models.UsrTournament{}, errs.ErrUsrTournamentNotFound
	}

	return usr, nil
}

func (u *UsrTournament) List(ctx context.Context, pars models.UsrTournamentListPars) ([]models.UsrTournament, int64, error) {
	return u.usrTourRepo.List(ctx, pars)
}

//func (u *UsrTournament) GetUsrTournaments(ctx context.Context, pars models.UsrGetTournamentsPars) ([]models.Tournament, int64, error) {
//return u.usrTourRepo.GetUsrTournaments(ctx, pars)
//}

func (u *UsrTournament) validateCU(ctx context.Context, obj models.UsrTournamentCU, id int64) error {
	forCreate := id == 0

	if forCreate && obj.UsrID == nil {
		return errs.ErrUsrIDNotFound
	}
	if forCreate && obj.TournamentID == nil {
		return errs.ErrTournamentIDNotFound
	}

	if forCreate && obj.UsrID != nil && obj.TournamentID != nil {
		_, err := u.Get(ctx, models.UsrTournamentGetPars{
			UsrID:        obj.UsrID,
			TournamentID: obj.TournamentID,
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

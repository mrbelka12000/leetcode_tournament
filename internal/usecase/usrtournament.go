package usecase

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

// UsrTournamentCreate ..
func (uc *UseCase) UsrTournamentCreate(ctx context.Context, obj models.UsrTournamentCU) (int64, error) {
	ses, err := uc.getSessionFromContext(ctx, true)
	if err != nil {
		return 0, err
	}

	obj.UsrID = &ses.UsrID

	return uc.cr.UsrTournament.Build(ctx, obj)
}

// UsrTournamentUpdate ..
func (uc *UseCase) UsrTournamentUpdate(ctx context.Context, obj models.UsrTournamentCU, id int64) error {
	ses, err := uc.getSessionFromContext(ctx, true)
	if err != nil {
		return err
	}

	if ses.TypeID != consts.UsrTypeAdmin {
		return errs.ErrPermissionDenied
	}

	err = uc.cr.UsrTournament.Update(ctx, obj, id)
	if err != nil {
		return fmt.Errorf("usr event update: %w", err)
	}

	return nil
}

// UsrTournamentGet ..
func (uc *UseCase) UsrTournamentGet(ctx context.Context, pars models.UsrTournamentGetPars) (models.UsrTournament, error) {
	return uc.cr.UsrTournament.Get(ctx, pars, true)
}

// UsrTournamentList ..
func (uc *UseCase) UsrTournamentList(ctx context.Context, pars models.UsrTournamentListPars) ([]models.UsrTournament, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.UsrTournament.List(ctx, pars)
}

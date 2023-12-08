package usecase

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

func (uc *UseCase) FillGeneral(ctx context.Context, data interface{}) (gen models.General) {
	gen.Data = data
	ses, err := uc.getSessionFromContext(ctx, true)
	if err != nil {
		return gen
	}

	gen.Usr, err = uc.cr.Usr.Get(ctx, models.UsrGetPars{
		ID: &ses.UsrID,
	}, true)
	if err != nil {
		return gen
	}

	return gen
}

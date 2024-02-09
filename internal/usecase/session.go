package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

func (uc *UseCase) getSessionFromContext(ctx context.Context, errNE bool) (models.Session, error) {
	token, ok := ctx.Value(consts.CKey).(string)
	if ok {
		ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
			Token: &token,
		})
		if err != nil {
			return models.Session{}, fmt.Errorf("get session by token: %w", err)
		}

		return ses, nil
	}
	if errNE {
		return models.Session{}, errors.New("not authorized")
	}

	return models.Session{}, nil
}

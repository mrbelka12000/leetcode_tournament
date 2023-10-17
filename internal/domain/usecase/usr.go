package usecase

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

func (uc *UseCase) UsrCreate(ctx context.Context, obj *models.UsrCU) (int64, error) {

	return uc.cr.Usr.Build(ctx, obj)
}

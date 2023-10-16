package interfaces

import (
	"context"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

type (
	Usr interface {
		Create(ctx context.Context, obj *models.UsrOld) (string, int64, error)
		Update(ctx context.Context, obj *models.UsrOld) error
		Get(ctx context.Context, id int64) (*models.UsrOld, error)
		List(ctx context.Context) ([]*models.UsrOld, int64, error)
		GetByField(ctx context.Context, field, value string) (*models.UsrOld, error)
	}

	Footprint interface {
		Create(ctx context.Context, obj *models.Footprint) (int64, error)
		Update(ctx context.Context, obj *models.Footprint, id int64) error
		Get(ctx context.Context, pars *models.FootprintGetPars) (*models.Footprint, error)
		List(ctx context.Context, pars *models.FootprintListPars) ([]*models.Footprint, int64, error)
	}
)

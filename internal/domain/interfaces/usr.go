package interfaces

import (
	"context"

	"leetcode_tournament/internal/domain/models"
)

type User interface {
	Create(ctx context.Context, obj *models.Usr) (string, error)
	Update(ctx context.Context, obj *models.Usr) error
	Get(ctx context.Context, id int64) (*models.Usr, error)
	List(ctx context.Context) ([]*models.Usr, int64, error)
	GetByField(ctx context.Context, field, value string) (*models.Usr, error)
}

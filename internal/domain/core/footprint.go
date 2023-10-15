package core

import (
	"context"
	"errors"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/interfaces"
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

type footprint struct {
	fpRepo interfaces.Footprint
}

func newFootprint(fpRepo interfaces.Footprint) *footprint {
	return &footprint{fpRepo: fpRepo}
}

func (f *footprint) Build(ctx context.Context, obj *models.Footprint) (int64, error) {
	err := f.validateCU(ctx, obj)
	if err != nil {
		return 0, err
	}

	return f.fpRepo.Create(ctx, obj)
}

func (f *footprint) Update(ctx context.Context, obj *models.Footprint, id int64) error {
	err := f.validateCU(ctx, obj)
	if err != nil {
		return err
	}

	return f.fpRepo.Update(ctx, obj, id)
}

func (f *footprint) Get(ctx context.Context, pars *models.FootprintGetPars) (*models.Footprint, error) {
	if pars.ID == nil && pars.UsrID == nil {
		return nil, errors.New("bad pars")
	}

	return f.fpRepo.Get(ctx, pars)
}

func (f *footprint) List(ctx context.Context, pars *models.FootprintListPars) ([]*models.Footprint, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, fmt.Errorf("invalid page size %v: %w", pars.PaginationParams.Limit, err)
	}

	return f.fpRepo.List(ctx, pars)
}

func (f *footprint) validateCU(ctx context.Context, obj *models.Footprint) error {
	if obj == nil {
		return errors.New("object is nil")
	}

	if obj.UsrID <= 0 {
		return errors.New("invalid usr id")
	}

	return nil
}

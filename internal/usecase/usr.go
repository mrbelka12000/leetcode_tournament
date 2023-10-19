package usecase

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

func (uc *UseCase) UsrCreate(ctx context.Context, obj models.UsrCU) (int64, string, error) {
	id, token, err := uc.cr.Usr.Build(ctx, obj)
	if err != nil {
		return 0, "", fmt.Errorf("usr build: %w", err)
	}

	err = uc.cr.Session.Build(ctx, models.Session{
		UsrID: id,
		Token: token,
	})
	if err != nil {
		return 0, "", fmt.Errorf("session build: %w", err)
	}

	return id, token, nil
}

func (uc *UseCase) Login(ctx context.Context, obj models.UsrLogin) (int64, string, error) {
	id, token, err := uc.cr.Usr.Login(ctx, obj)
	if err != nil {
		return 0, "", fmt.Errorf("usr login: %w", err)
	}

	err = uc.cr.Session.Build(ctx, models.Session{
		UsrID: id,
		Token: token,
	})
	if err != nil {
		return 0, "", fmt.Errorf("session build: %w", err)
	}

	return id, token, nil
}

package usecase

import (
	"context"
	"fmt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

func (uc *UseCase) Registration(ctx context.Context, obj models.UsrCU) (int64, string, error) {
	ctx, err := uc.cr.Transaction.BeginTransaction(ctx)
	if err != nil {
		return 0, "", err
	}

	defer uc.cr.Transaction.RollbackTransaction(ctx)

	id, token, err := uc.cr.Usr.Build(ctx, obj)
	if err != nil {
		return 0, "", fmt.Errorf("usr build: %w", err)
	}

	leetCodeResp, err := uc.cr.Score.Stats(ctx, *obj.Name)
	if err != nil {
		return 0, "", err
	}

	id, err = uc.cr.Score.Build(ctx, models.ScoreCU{
		UsrID: &id,
		Current: &models.Points{
			Easy:   leetCodeResp.Easy,
			Medium: leetCodeResp.Medium,
			Hard:   leetCodeResp.Hard,
			Total:  leetCodeResp.Total,
		},
	})
	if err != nil {
		return 0, "", err
	}

	fmt.Println(id, token)
	err = uc.cr.Session.Build(ctx, models.Session{
		UsrID:  id,
		Token:  token,
		TypeID: *obj.TypeID,
	})
	if err != nil {
		return 0, "", fmt.Errorf("session build: %w", err)
	}

	err = uc.cr.Transaction.CommitTransaction(ctx)
	if err != nil {
		return 0, "", err
	}

	return id, token, nil
}

func (uc *UseCase) Login(ctx context.Context, obj models.UsrLogin) (int64, string, error) {
	id, token, err := uc.cr.Usr.Login(ctx, obj)
	if err != nil {
		return 0, "", fmt.Errorf("usr login: %w", err)
	}

	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		UsrID: &id,
	})
	if err != nil {
		usr, err := uc.cr.Usr.Get(ctx, models.UsrGetPars{
			ID: &id,
		}, true)
		if err != nil {
			return 0, "", fmt.Errorf("get usr: %w", err)
		}

		err = uc.cr.Session.Build(ctx, models.Session{
			UsrID:  id,
			Token:  token,
			TypeID: usr.TypeID,
		})
		if err != nil {
			return 0, "", fmt.Errorf("session build: %w", err)
		}

		return id, token, nil
	}

	return id, ses.Token, nil
}

func (uc *UseCase) UsrUpdate(ctx context.Context, obj models.UsrCU) error {
	token := ctx.Value(consts.CKey).(string)

	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return fmt.Errorf("get session by token: %w", err)
	}

	return uc.cr.Usr.Update(ctx, obj, ses.UsrID)
}

func (uc *UseCase) UsrList(ctx context.Context, pars models.UsrListPars) ([]models.Usr, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	return uc.cr.Usr.List(ctx, pars)
}

func (uc *UseCase) UsrProfile(ctx context.Context) (models.Usr, error) {
	token := ctx.Value(consts.CKey).(string)
	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return models.Usr{}, fmt.Errorf("get session by token: %w", err)
	}

	return uc.cr.Usr.Get(ctx, models.UsrGetPars{
		ID: &ses.ID,
	}, true)
}

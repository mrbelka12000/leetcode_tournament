package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

// Registration ..
func (uc *UseCase) Registration(ctx context.Context, obj models.UsrCU) (int64, string, error) {
	ctx, err := uc.cr.Tx.ContextWithTransaction(ctx)
	if err != nil {
		return 0, "", err
	}
	defer uc.cr.Tx.RollbackContextTransaction(ctx)

	id, token, err := uc.cr.Usr.Build(ctx, obj)
	if err != nil {
		return 0, "", fmt.Errorf("usr build: %w", err)
	}

	stats, err := uc.cr.Score.Stats(ctx, *obj.Username)
	if err != nil {
		return 0, "", fmt.Errorf("get usr stats: %w", err)
	}

	point := models.Points(stats)
	_, err = uc.cr.Score.Build(ctx, models.ScoreCU{
		UsrID:   &id,
		Current: &point,
	})
	if err != nil {
		return 0, "", fmt.Errorf("score build: %w", err)
	}

	err = uc.cr.Session.Build(ctx, models.Session{
		UsrID:  id,
		Token:  token,
		TypeID: *obj.TypeID,
	})
	if err != nil {
		return 0, "", fmt.Errorf("session build: %w", err)
	}

	err = uc.cr.Tx.CommitContextTransaction(ctx)
	if err != nil {
		return 0, "", err
	}

	return id, token, nil
}

// Login ..
func (uc *UseCase) Login(ctx context.Context, obj models.UsrLogin) (int64, string, error) {
	ctx, err := uc.cr.Tx.ContextWithTransaction(ctx)
	if err != nil {
		return 0, "", err
	}
	defer uc.cr.Tx.RollbackContextTransaction(ctx)

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

	err = uc.cr.Tx.CommitContextTransaction(ctx)
	if err != nil {
		return 0, "", err
	}

	return id, ses.Token, nil
}

// UsrUpdate ..
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

// UsrList ..
func (uc *UseCase) UsrList(ctx context.Context, pars models.UsrListPars) ([]models.Usr, int64, error) {
	if err := validator.RequirePageSize(pars.PaginationParams); err != nil {
		return nil, 0, err
	}

	result, tCount, err := uc.cr.Usr.List(ctx, pars)
	if err != nil {
		return nil, 0, fmt.Errorf("get usrs list: %w", err)
	}

	for i := 0; i < len(result); i++ {
		result[i].Score, err = uc.cr.Score.Get(ctx, models.ScoreGetPars{
			UsrID: &result[i].ID,
		}, true)
		if err != nil {
			return nil, 0, fmt.Errorf("get score: %w", err)
		}
	}

	return result, tCount, nil
}

// UsrProfile ..
func (uc *UseCase) UsrProfile(ctx context.Context) (models.Usr, error) {
	token := ctx.Value(consts.CKey).(string)
	ses, err := uc.cr.Session.Get(ctx, models.SessionGetPars{
		Token: &token,
	})
	if err != nil {
		return models.Usr{}, fmt.Errorf("get session by token: %w", err)
	}

	return uc.UsrGet(ctx, models.UsrGetPars{
		ID:        &ses.ID,
		WithScore: true,
	})
}

func (uc *UseCase) UsrGet(ctx context.Context, pars models.UsrGetPars) (models.Usr, error) {
	usr, err := uc.cr.Usr.Get(ctx, pars, true)
	if err != nil {
		return models.Usr{}, fmt.Errorf("get usr: %w", err)
	}

	if pars.WithScore {
		usr.Score, err = uc.cr.Score.Get(ctx, models.ScoreGetPars{
			UsrID: &usr.ID,
		}, true)
		if err != nil {
			return models.Usr{}, fmt.Errorf("get score for usr: %w", err)
		}
	}

	return usr, nil
}

func (uc *UseCase) UsrScoreUpdate(ctx context.Context) error {
	ctx, err := uc.cr.Tx.ContextWithTransaction(ctx)
	if err != nil {
		return err
	}
	defer uc.cr.Tx.RollbackContextTransaction(ctx)

	usrs, _, err := uc.cr.Usr.List(ctx, models.UsrListPars{})
	if err != nil {
		return fmt.Errorf("usr list: %w", err)
	}

	for i := 0; i < len(usrs); i++ {
		time.Sleep(2 * time.Second)
		score, err := uc.cr.Score.Get(ctx, models.ScoreGetPars{
			UsrID: &usrs[i].ID,
		}, true)
		if err != nil {
			return fmt.Errorf("score get: %w", err)
		}

		stats, err := uc.cr.Score.Stats(ctx, usrs[i].Username)
		if err != nil {
			log.Printf("can not get stats from leetocode for usr: %v \n", usrs[i].Username)
			continue
		}

		point := models.Points(stats)

		err = uc.cr.Score.Update(ctx, models.ScoreCU{
			Current: &point,
		}, score.ID)
		if err != nil {
			return fmt.Errorf("score update: %w", err)
		}
	}

	err = uc.cr.Tx.CommitContextTransaction(ctx)
	if err != nil {
		return err
	}

	uc.log.Info().Msg("UsrScoreUpdate successfully worked")
	return nil
}

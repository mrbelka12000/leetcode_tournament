package core

import (
	"context"
	"errors"
	"fmt"

	"leetcode_tournament/internal/domain/interfaces"
	"leetcode_tournament/internal/domain/models"
	"leetcode_tournament/pkg/generator"
)

type usr struct {
	usrRepo       interfaces.User
	leetCodeStats interfaces.LeetCodeStats
}

func newUsr(usrRepo interfaces.User, leetCodeStats interfaces.LeetCodeStats) *usr {
	return &usr{
		usrRepo:       usrRepo,
		leetCodeStats: leetCodeStats,
	}
}

func (u *usr) Create(ctx context.Context, obj *models.Usr) (string, error) {
	err := u.validateCU(ctx, obj)
	if err != nil {
		return "", fmt.Errorf("validate CU: %w", err)
	}

	obj.Secret = generator.RandString(8)

	return u.usrRepo.Create(ctx, obj)
}

func (u *usr) Update(ctx context.Context, obj *models.Usr) error {
	err := u.validateCU(ctx, obj)
	if err != nil {
		return err
	}

	return u.usrRepo.Update(ctx, obj)
}

func (u *usr) Get(ctx context.Context, id int64) (*models.Usr, error) {
	return u.usrRepo.Get(ctx, id)
}

func (u *usr) List(ctx context.Context) ([]*models.Usr, int64, error) {
	return u.usrRepo.List(ctx)
}

func (u *usr) validateCU(ctx context.Context, obj *models.Usr) error {

	forCreate := obj.Secret == ""

	var (
		origUsr *models.Usr
		err     error
	)
	if !forCreate {
		origUsr, err = u.usrRepo.GetByField(ctx, "secret", obj.Secret)
		if err != nil {
			return err
		}
	}
	if obj.Nickname == "" {
		return errors.New("empty nickname")
	}

	if obj.Name == "" {
		return errors.New("empty name")
	}

	if !forCreate {
		usr, err := u.usrRepo.GetByField(ctx, "nickname", obj.Nickname)
		if err == nil && origUsr.ID != usr.ID {
			return errors.New("nickname already exists")
		}
	}
	return nil
}

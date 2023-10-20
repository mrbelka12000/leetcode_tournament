package usr

import (
	"context"
	"fmt"

	"github.com/AlekSi/pointer"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/errs"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/pkg/ptr"
	"github.com/mrbelka12000/leetcode_tournament/pkg/validator"
)

type Usr struct {
	usrRepo       Repo
	leetCodeStats LeetCodeStats
}

func New(usrRepo Repo, leetCodeStats LeetCodeStats) *Usr {
	return &Usr{
		usrRepo:       usrRepo,
		leetCodeStats: leetCodeStats,
	}
}

func (u *Usr) Build(ctx context.Context, obj models.UsrCU) (int64, string, error) {
	obj.StatusID = ptr.UsrStatusPointer(consts.UsrStatusCreated)
	err := u.validateCU(ctx, obj, 0)
	if err != nil {
		return 0, "", fmt.Errorf("validate CU: %w", err)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*obj.Password), 14)
	if err != nil {
		return 0, "", fmt.Errorf("generate password hash: %w", err)
	}

	*obj.Password = string(hashedPassword)

	id, err := u.usrRepo.Create(ctx, obj)
	if err != nil {
		return 0, "", fmt.Errorf("create usr: %w", err)
	}

	return id, uuid.New().String(), nil
}

func (u *Usr) Login(ctx context.Context, obj models.UsrLogin) (int64, string, error) {
	usr, err := u.Get(ctx, models.UsrGetPars{
		UsernameEmail: &obj.UsernameEmail,
	}, true)
	if err != nil {
		return 0, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(obj.Password))
	if err != nil {
		return 0, "", errs.ErrPasswordDontMatch
	}

	return usr.ID, uuid.New().String(), nil
}

func (u *Usr) Update(ctx context.Context, obj models.UsrCU, id int64) error {
	err := u.validateCU(ctx, obj, id)
	if err != nil {
		return err
	}

	return u.usrRepo.Update(ctx, obj, id)
}

func (u *Usr) Get(ctx context.Context, pars models.UsrGetPars, errNE bool) (models.Usr, error) {
	usr, err := u.usrRepo.Get(ctx, pars)
	if err != nil {
		return models.Usr{}, fmt.Errorf("usr get from db: %w", err)
	}
	if errNE && usr.ID == 0 {
		return models.Usr{}, errs.ErrUsrStatusNotFound
	}

	return usr, nil
}

func (u *Usr) List(ctx context.Context, pars models.UsrListPars) ([]models.Usr, int64, error) {
	return u.usrRepo.List(ctx, pars)
}

func (u *Usr) validateCU(ctx context.Context, obj models.UsrCU, id int64) error {

	forCreate := id == 0

	if forCreate && obj.Username == nil {
		return errs.ErrUsernameNotFound
	}
	if obj.Username != nil {
		if len([]rune(*obj.Username)) >= 255 || len([]rune(*obj.Username)) == 0 {
			return errs.ErrInvalidUsername
		}
		usr, err := u.Get(ctx, models.UsrGetPars{
			Username: pointer.ToString(*obj.Username),
		}, true)
		if err == nil {
			if usr.ID != id {
				return errs.ErrUsernameExists
			}
		}

	}

	if forCreate && obj.Email == nil {
		return errs.ErrEmailNotFound
	}
	if obj.Email != nil {
		err := validator.ValidateEmail(*obj.Email)
		if err != nil {
			return errs.ErrInvalidEmail
		}
		usr, err := u.Get(ctx, models.UsrGetPars{
			Email: pointer.ToString(*obj.Email),
		}, true)
		if err == nil {
			if usr.ID != id {
				return errs.ErrEmailExists
			}
		}
	}

	if forCreate && obj.Name == nil {
		return errs.ErrNameNotFound
	}
	if obj.Name != nil {
		if len([]rune(*obj.Name)) >= 255 || len([]rune(*obj.Name)) == 0 {
			return errs.ErrInvalidUsername
		}
	}

	if forCreate && obj.Password == nil {
		return errs.ErrPasswordNotFound
	}
	if forCreate && obj.Password != nil {
		err := validator.ValidatePassword(*obj.Password)
		if err != nil {
			return err
		}
	}

	if forCreate && obj.StatusID == nil {
		return errs.ErrUsrStatusNotFound
	}
	if obj.StatusID != nil {
		if !consts.IsValidUsrStatus(*obj.StatusID) {
			return errs.ErrInvalidUsrStatus
		}
	}

	if forCreate && obj.TypeID == nil {
		return errs.ErrUsrTypeNotFound
	}
	if obj.TypeID != nil {
		if !consts.IsValidUsrType(*obj.TypeID) {
			return errs.ErrInvalidUsrType
		}
	}

	return nil
}

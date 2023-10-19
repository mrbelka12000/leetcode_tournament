package validator

import (
	"errors"
	"net/mail"
	"unicode"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

func RequirePageSize(pars models.PaginationParams) error {
	if pars.Limit <= 0 {
		return errors.New("incorrect page size")
	}

	return nil
}

func ValidatePassword(password string) error {
	if len([]rune(password)) < 8 {
		return errors.New("minimum len of password is 8")
	}
	var haveLetter, haveNumber bool

	for _, digit := range password {
		if unicode.IsLetter(digit) {
			haveLetter = true
		}
		if unicode.IsNumber(digit) {
			haveNumber = true
		}
	}

	if !haveNumber {
		return errors.New("password must have one number")
	}
	if !haveLetter {
		return errors.New("password must have one letter")
	}

	return nil
}

func ValidateEmail(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

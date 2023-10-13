package validator

import (
	"errors"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

func RequirePageSize(pars models.PaginationParams) error {
	if pars.Limit == 0 {
		return errors.New("incorrect page size")
	}

	return nil
}

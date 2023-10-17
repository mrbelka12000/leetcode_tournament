package usecase

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/service"
)

type UseCase struct {
	cr *service.Core
}

func New(cr *service.Core) *UseCase {
	return &UseCase{cr: cr}
}

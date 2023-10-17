package usecase

import "github.com/mrbelka12000/leetcode_tournament/internal/domain/core"

type UseCase struct {
	cr *core.Core
}

func New(cr *core.Core) *UseCase {
	return &UseCase{cr: cr}
}

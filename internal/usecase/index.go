package usecase

import (
	"github.com/rs/zerolog"

	"github.com/mrbelka12000/leetcode_tournament/internal/service"
)

type UseCase struct {
	cr  *service.Core
	log zerolog.Logger
}

func New(cr *service.Core, log zerolog.Logger) *UseCase {
	return &UseCase{
		cr:  cr,
		log: log,
	}
}

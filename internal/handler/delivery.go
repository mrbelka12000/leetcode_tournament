package handler

import (
	"github.com/gorilla/schema"

	"github.com/mrbelka12000/leetcode_tournament/internal/usecase"
)

const (
	templateDir = "templates/"
)

type Handler struct {
	uc      *usecase.UseCase
	decoder *schema.Decoder
}

func New(uc *usecase.UseCase) *Handler {
	return &Handler{
		uc:      uc,
		decoder: schema.NewDecoder(),
	}
}

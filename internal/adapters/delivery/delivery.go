package delivery

import (
	"github.com/gorilla/schema"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/usecase"
)

const (
	templateDir = "templates/"
)

type DeliveryHTTP struct {
	uc      *usecase.UseCase
	decoder *schema.Decoder
}

func NewDeliveryHTTP(uc *usecase.UseCase) *DeliveryHTTP {
	return &DeliveryHTTP{
		uc:      uc,
		decoder: schema.NewDecoder(),
	}
}

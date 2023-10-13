package delivery

import (
	"github.com/mrbelka12000/leetcode_tournament/internal/domain/core"
)

const (
	templateDir = "templates/"
)

type DeliveryHTTP struct {
	cr *core.Core
}

func NewDeliveryHTTP(cr *core.Core) *DeliveryHTTP {
	return &DeliveryHTTP{cr: cr}
}

package handler

import (
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
	"github.com/rs/zerolog"

	"github.com/mrbelka12000/leetcode_tournament/internal/usecase"
)

type Handler struct {
	uc      *usecase.UseCase
	decoder *schema.Decoder
	limiter rateLimit
	log     zerolog.Logger
}

func New(uc *usecase.UseCase, limiter rateLimit, log zerolog.Logger) *Handler {
	return &Handler{
		uc:      uc,
		decoder: schema.NewDecoder(),
		limiter: limiter,
		log:     log,
	}
}

func (h *Handler) uExtractPaginationPars(pars url.Values) (offset int64, limit int64, page int64) {
	var err error

	qPar := pars.Get("page_size")
	if qPar != "" {
		limit, err = strconv.ParseInt(qPar, 10, 64)
		if err != nil {
			limit = 0
		}
	}

	qPar = pars.Get("page")
	if qPar != "" {
		page, err = strconv.ParseInt(qPar, 10, 64)
		if err != nil {
			page = 0
		}
	}
	if page == 0 {
		page = 1
	}

	offset = (page - 1) * limit

	return offset, limit, page
}

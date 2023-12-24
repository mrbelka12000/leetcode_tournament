package handler

import (
	"context"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/internal/view/pages"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	var pars models.UsrListPars

	pars.Limit = 20
	var page int64 = 1

	usrs, tCount, err := h.uc.UsrList(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := h.uc.FillGeneral(
		r.Context(),
		usrs,
	)

	component := pages.Index(data, page, pars.Limit, tCount)
	component.Render(context.Background(), w)
}

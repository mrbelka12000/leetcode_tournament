package handler

import (
	"context"
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/gorilla/mux"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/internal/view/pages"
)

// EventCreate ..
func (h *Handler) EventCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.EventCU
	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.uc.EventCreate(r.Context(), obj)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// alert := consts.SuccessAlert{
	// 	AlertType:    consts.Success,
	// 	AlertMessage: "Event successfully created",
	// }

	// view.RenderTemplate(w, r, "alert", alert)
}

// EventUpdate ..
func (h *Handler) EventUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.EventCU
	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.log.Err(err).Send()
		http.Error(w, "no id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.EventUpdate(r.Context(), obj, id)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// alert := consts.SuccessAlert{
	// 	AlertType:    consts.Success,
	// 	AlertMessage: "Event successfully updated",
	// }

	// view.RenderTemplate(w, r, "alert", alert)
}

// EventGet ..
func (h *Handler) EventGet(w http.ResponseWriter, r *http.Request) {

	var pars models.EventGetPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		h.log.Err(err).Send()
		http.Error(w, "no id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	pars.ID = pointer.ToInt64(id)
	eventPage, err := h.uc.EventPageGet(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := pages.EventDetails(h.uc.FillGeneral(
		r.Context(),
		eventPage,
	))

	component.Render(context.Background(), w)
}

func (h *Handler) EventList(w http.ResponseWriter, r *http.Request) {
	var pars models.EventListPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var page int64
	pars.Offset, pars.Limit, page = h.uExtractPaginationPars(r.URL.Query())

	events, tCount, err := h.uc.EventList(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	component := pages.EventsPage(h.uc.FillGeneral(
		r.Context(),
		events,
	), page, pars.Limit, tCount)

	component.Render(context.Background(), w)
}

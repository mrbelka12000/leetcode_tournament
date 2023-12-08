package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/gorilla/mux"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

// UsrTournamentCreate ..
func (h *Handler) UsrTournamentCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.UsrTournamentCU
	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.uc.UsrTournamentCreate(r.Context(), obj)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UsrTournamentUpdate ..
func (h *Handler) UsrTournamentUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.UsrTournamentCU
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

	err = h.uc.UsrTournamentUpdate(r.Context(), obj, id)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UsrTournamentGet ..
func (h *Handler) UsrTournamentGet(w http.ResponseWriter, r *http.Request) {

	var pars models.UsrTournamentGetPars
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

	usrTournament, err := h.uc.UsrTournamentGet(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v \n", usrTournament)
}

func (h *Handler) UsrTournamentList(w http.ResponseWriter, r *http.Request) {
	var pars models.UsrTournamentListPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var page int64
	pars.Offset, pars.Limit, page = h.uExtractPaginationPars(r.URL.Query())

	events, tCount, err := h.uc.UsrTournamentList(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.PaginatedListRepSt{
		Page:       page,
		PageSize:   pars.Limit,
		TotalCount: tCount,
		Results:    events,
	}

	fmt.Println(resp)
}

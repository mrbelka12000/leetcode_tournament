package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/gorilla/mux"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

// TournamentCreate ..
func (h *Handler) TournamentCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.TournamentCU
	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.uc.TournamentCreate(r.Context(), obj)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TournamentUpdate ..
func (h *Handler) TournamentUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.TournamentCU
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

	err = h.uc.TournamentUpdate(r.Context(), obj, id)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// TournamentGet ..
func (h *Handler) TournamentGet(w http.ResponseWriter, r *http.Request) {

	var pars models.TournamentGetPars
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

	tournament, err := h.uc.TournamentGet(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v \n", tournament)
}

func (h *Handler) TournamentList(w http.ResponseWriter, r *http.Request) {
	var pars models.TournamentListPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var page int64
	pars.Offset, pars.Limit, page = h.uExtractPaginationPars(r.URL.Query())

	tournaments, tCount, err := h.uc.TournamentList(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.PaginatedListRepSt{
		Page:       page,
		PageSize:   pars.Limit,
		TotalCount: tCount,
		Results:    tournaments,
	}

	fmt.Println(resp)
}

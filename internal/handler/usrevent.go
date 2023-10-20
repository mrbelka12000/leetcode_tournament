package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AlekSi/pointer"
	"github.com/gorilla/mux"

	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

// UsrEventCreate ..
func (h *Handler) UsrEventCreate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.UsrEventCU
	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = h.uc.UsrEventCreate(r.Context(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UsrEventUpdate ..
func (h *Handler) UsrEventUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.UsrEventCU
	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "no id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.UsrEventUpdate(r.Context(), obj, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

// UsrEventGet ..
func (h *Handler) UsrEventGet(w http.ResponseWriter, r *http.Request) {

	var pars models.UsrEventGetPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vars := mux.Vars(r)
	idStr, ok := vars["id"]
	if !ok {
		http.Error(w, "no id in path", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	pars.ID = pointer.ToInt64(id)

	event, err := h.uc.UsrEventGet(r.Context(), pars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("%+v \n", event)
}

func (h *Handler) UsrEventList(w http.ResponseWriter, r *http.Request) {
	var pars models.UsrEventListPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var page int64
	pars.Offset, pars.Limit, page = h.uExtractPaginationPars(r.URL.Query())

	events, tCount, err := h.uc.UsrEventList(r.Context(), pars)
	if err != nil {
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

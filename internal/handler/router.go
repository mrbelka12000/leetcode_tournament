package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Index).Methods(http.MethodGet)
	r.HandleFunc("/users", h.GetAllUsers).Methods(http.MethodGet)
	r.HandleFunc("/registration", h.Registration).Methods(http.MethodPost)
	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	r.HandleFunc("/usr_update", h.UsrUpdate).Methods(http.MethodPost)
	return r
}

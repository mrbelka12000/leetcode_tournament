package handler

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (h *Handler) InitRoutes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", h.Index).Methods(http.MethodGet)

	r.HandleFunc("/registration", h.Registration).Methods(http.MethodPost)
	r.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	r.HandleFunc("/usr/update", h.getCookie(h.ProfileUpdate)).Methods(http.MethodPost)
	r.HandleFunc("/usr", h.getCookie(h.GetUsr)).Methods(http.MethodGet)
	r.HandleFunc("/users", h.Usrs).Methods(http.MethodGet)

	r.HandleFunc("/event", h.getCookie(h.EventCreate)).Methods(http.MethodPost)
	r.HandleFunc("/event/update/{id}", h.getCookie(h.EventUpdate)).Methods(http.MethodPost)
	r.HandleFunc("/event", h.EventList).Methods(http.MethodGet)
	r.HandleFunc("/event/{id}", h.EventGet).Methods(http.MethodGet)

	r.HandleFunc("/usr_event", h.getCookie(h.UsrEventCreate)).Methods(http.MethodPost)
	r.HandleFunc("/usr_event/update/{id}", h.getCookie(h.UsrEventUpdate)).Methods(http.MethodPost)
	r.HandleFunc("/usr_event", h.UsrEventList).Methods(http.MethodGet)
	r.HandleFunc("/usr_event/{id}", h.UsrEventGet).Methods(http.MethodGet)

	r.HandleFunc("/tournament", h.getCookie(h.TournamentCreate)).Methods(http.MethodPost)
	r.HandleFunc("/tournament/update/{id}", h.getCookie(h.TournamentUpdate)).Methods(http.MethodPost)
	r.HandleFunc("/tournament", h.TournamentList).Methods(http.MethodGet)
	r.HandleFunc("/tournament/{id}", h.TournamentGet).Methods(http.MethodGet)

	r.HandleFunc("/usr_event", h.getCookie(h.UsrEventCreate)).Methods(http.MethodPost)
	r.HandleFunc("/usr_event/update/{id}", h.getCookie(h.UsrEventUpdate)).Methods(http.MethodPost)
	r.HandleFunc("/usr_event", h.UsrEventList).Methods(http.MethodGet)
	r.HandleFunc("/usr_event/{id}", h.UsrEventGet).Methods(http.MethodGet)

	r.HandleFunc("/usr_tournament", h.getCookie(h.UsrTournamentCreate)).Methods(http.MethodPost)
	r.HandleFunc("/usr_tournament/update/{id}", h.getCookie(h.UsrTournamentUpdate)).Methods(http.MethodPost)
	r.HandleFunc("/usr_tournament", h.UsrTournamentList).Methods(http.MethodGet)
	r.HandleFunc("/usr_tournament/{id}", h.UsrTournamentGet).Methods(http.MethodGet)

	r.Use(h.limit)

	return r
}

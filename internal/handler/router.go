package handler

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
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

	return r
}

func (h *Handler) getCookie(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie(consts.CookieName)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), consts.CKey, cookie.Value)
		next(w, r.WithContext(ctx))
		return
	}
}

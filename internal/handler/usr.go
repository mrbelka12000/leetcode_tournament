package handler

import (
	"context"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
	"github.com/mrbelka12000/leetcode_tournament/internal/view/components"
)

// Registration ..
func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usr models.UsrCU

	err = h.decoder.Decode(&usr, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, token, err := h.uc.Registration(r.Context(), usr)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cookie := http.Cookie{
		Name:     consts.CookieName,
		Value:    token,
		Path:     "*",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Add("Hx-trigger", "usersUpdate")

	alert := consts.SuccessAlert{
		AlertType:      consts.Success,
		AlertMessage:   "Successfully registered",
		ButtonIdToHide: "registerButton",
	}

	component := components.Alert(alert)
	component.Render(context.Background(), w)
}

// Login ..
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usrLogin models.UsrLogin
	err = h.decoder.Decode(&usrLogin, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, token, err := h.uc.Login(r.Context(), usrLogin)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, "User not found or incorrect password", http.StatusBadRequest)
		return
	}
	cookie := http.Cookie{
		Name:     "session",
		Value:    token,
		Path:     "*",
		MaxAge:   3600,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, &cookie)

	w.Header().Add("Hx-trigger", "headerUpdate")

	alert := consts.SuccessAlert{
		AlertType:      consts.Success,
		AlertMessage:   "Successfully logged in",
		ButtonIdToHide: "loginButton",
	}

	component := components.Alert(alert)
	component.Render(context.Background(), w)
}

func (h *Handler) Header(w http.ResponseWriter, r *http.Request) {
	data := h.uc.FillGeneral(
		r.Context(),
		nil,
	)
	component := components.Header(data.Usr.Name)
	component.Render(context.Background(), w)
}

// ProfileUpdate ..
func (h *Handler) ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.UsrCU

	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.UsrUpdate(r.Context(), obj)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Hx-trigger", "usersUpdate")

	alert := consts.SuccessAlert{
		AlertType:      consts.Success,
		AlertMessage:   "Successfully updated",
		ButtonIdToHide: "updateUserButton",
	}

	component := components.Alert(alert)
	component.Render(context.Background(), w)
}

// Usrs ..
func (h *Handler) Usrs(w http.ResponseWriter, r *http.Request) {
	var pars models.UsrListPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var page int64
	pars.Offset, pars.Limit, page = h.uExtractPaginationPars(r.URL.Query())

	usrs, tCount, err := h.uc.UsrList(r.Context(), pars)
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	component := components.Users(h.uc.FillGeneral(
		r.Context(),
		usrs,
	), page, pars.Limit, tCount)
	component.Render(context.Background(), w)
}

func (h *Handler) GetUsr(w http.ResponseWriter, r *http.Request) {
	usr, err := h.uc.UsrProfile(r.Context())
	if err != nil {
		h.log.Err(err).Send()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = usr
}

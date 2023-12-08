package handler

import (
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
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

	RenderTemplate(w, "alert", alert)
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

	alert := consts.SuccessAlert{
		AlertType:      consts.Success,
		AlertMessage:   "Successfully logged in",
		ButtonIdToHide: "loginButton",
	}

	RenderTemplate(w, "alert", alert)
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

	RenderTemplate(w, "update", nil)
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

	RenderTemplate(w, "users-table", models.PaginatedListRepSt{
		Page:       page,
		PageSize:   pars.Limit,
		TotalCount: tCount,
		Results: h.uc.FillGeneral(
			r.Context(),
			usrs,
		),
	})
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

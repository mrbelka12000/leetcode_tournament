package handler

import (
	"html/template"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
	"github.com/mrbelka12000/leetcode_tournament/internal/models"
)

// Registration ..
func (h *Handler) Registration(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usr models.UsrCU

	err = h.decoder.Decode(&usr, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, token, err := h.uc.Registration(r.Context(), usr)
	if err != nil {
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

	renderTemplate(w, "register", nil)
}

// Login ..
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var usrLogin models.UsrLogin
	err = h.decoder.Decode(&usrLogin, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, token, err := h.uc.Login(r.Context(), usrLogin)
	if err != nil {
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

	renderTemplate(w, "login", nil)
}

// ProfileUpdate ..
func (h *Handler) ProfileUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var obj models.UsrCU

	err = h.decoder.Decode(&obj, r.Form)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.uc.UsrUpdate(r.Context(), obj)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Add("Hx-trigger", "usersUpdate")

	renderTemplate(w, "update", nil)
}

// Usrs ..
func (h *Handler) Usrs(w http.ResponseWriter, r *http.Request) {

	var pars models.UsrListPars
	err := h.decoder.Decode(&pars, r.URL.Query())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var page int64
	pars.Offset, pars.Limit, page = h.uExtractPaginationPars(r.URL.Query())

	usrs, tCount, err := h.uc.UsrList(r.Context(), pars)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	renderTemplate(w, "users-table", models.PaginatedListRepSt{
		Page:       page,
		PageSize:   pars.Limit,
		TotalCount: tCount,
		Results:    usrs,
	})
}

func (h *Handler) GetUsr(w http.ResponseWriter, r *http.Request) {

	usr, err := h.uc.UsrProfile(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_ = usr
}

func renderTemplate(w http.ResponseWriter, templateName string, result interface{}) {
	t, err := template.ParseFiles(templateDir + templateName + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

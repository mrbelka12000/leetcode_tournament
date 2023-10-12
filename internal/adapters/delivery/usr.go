package delivery

import (
	"net/http"
	"text/template"

	"leetcode_tournament/internal/domain/models"
)

func (d *DeliveryHTTP) CreateUser(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.Usr{
		Name:     r.FormValue("name"),
		Nickname: r.FormValue("nickname"),
	}

	_, err = d.cr.Usr.Create(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := template.ParseFiles(templateDir + "register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, user.Secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d *DeliveryHTTP) UpdateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.Usr{
		Name:     r.FormValue("name"),
		Nickname: r.FormValue("nickname"),
		Secret:   r.FormValue("secret"),
	}

	err = d.cr.Usr.Update(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

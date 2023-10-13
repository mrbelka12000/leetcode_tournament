package delivery

import (
	"net/http"
	"text/template"

	"github.com/mrbelka12000/leetcode_tournament/internal/domain/models"
)

func (d *DeliveryHTTP) UsrCreate(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user := &models.Usr{
		Name:     r.FormValue("name"),
		Nickname: r.FormValue("nickname"),
	}

	secret, _, err := d.cr.Usr.Create(r.Context(), user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	t, err := template.ParseFiles(templateDir + "register.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, secret)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (d *DeliveryHTTP) UsrUpdate(w http.ResponseWriter, r *http.Request) {
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

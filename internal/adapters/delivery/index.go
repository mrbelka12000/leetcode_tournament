package delivery

import (
	"net/http"
	"text/template"
)

func (d *DeliveryHTTP) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templateDir + "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	index, err := d.cr.MainPage(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, index)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

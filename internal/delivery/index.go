package delivery

import (
	"html/template"
	"net/http"
)

func (d *DeliveryHTTP) Index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles(templateDir + "index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

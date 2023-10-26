package handler

import (
	"html/template"
	"net/http"

	"github.com/mrbelka12000/leetcode_tournament/internal/consts"
)

func RenderTemplate(w http.ResponseWriter, templateName string, result interface{}) {
	t, err := template.ParseFiles(consts.TemplateDir + templateName + ".html")
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

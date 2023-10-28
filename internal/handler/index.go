package handler

import (
	"net/http"
)

func (h *Handler) Index(w http.ResponseWriter, r *http.Request) {
	RenderTemplate(w, "index", nil)
}

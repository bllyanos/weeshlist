package handlers

import (
	"net/http"

	"github.com/bllyanos/weeshlist/internal/pkg/comm"
	"github.com/bllyanos/weeshlist/internal/pkg/comm/render"
	"github.com/bllyanos/weeshlist/internal/pkg/tp"
)

type SearchHandler struct {
	templates *tp.Template
}

func NewSearchHandler(templates *tp.Template) *SearchHandler {
	return &SearchHandler{templates: templates}
}

func (h *SearchHandler) SearchHandler(w http.ResponseWriter, r *http.Request) {
	// get query parameters
	s := r.URL.Query().Get("s")

	// render search page
	render.Page(
		h.templates,
		w,
		r,
		"search",
		nil, comm.JSON{
			"Search": s,
		})
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/bllyanos/weeshlist/internal/pkg/comm"
	"github.com/bllyanos/weeshlist/internal/pkg/comm/render"
	"github.com/bllyanos/weeshlist/internal/pkg/tp"
)

type IndexHandler struct {
	templates *tp.Template
}

func NewIndexHandler(templates *tp.Template) *IndexHandler {
	handler := &IndexHandler{
		templates: templates,
	}
	return handler
}

func (h *IndexHandler) IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Println("Masuk sini banget")
		render.NotFound(h.templates, w, r)
		return
	}
	data := comm.JSON{
		"Wishes": []comm.JSON{
			{"Name": "Wish 1", "Description": "lorem ipsum bla bla bla bla bla", "Price": 100_000, "URL": "https://google.com/"},
			{"Name": "Wish 2", "Description": "lorem ipsum bla bla bla bla bla", "Price": 200_000},
			{"Name": "Wish 3", "Description": "lorem ipsum bla bla bla bla bla", "Price": 300_000},
			{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
			{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
			{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
			{"Name": "Wish 4", "Description": "lorem ipsum bla bla bla bla bla", "Price": 400_000},
		},
	}
	render.Page(h.templates, w, r, "index", nil, data)
}

// 404 handler
func (h *IndexHandler) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Masuk sini")
	render.NotFound(h.templates, w, r)
}

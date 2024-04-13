package handlers

import (
	"fmt"
	"net/http"

	"github.com/bllyanos/weeshlist/internal/pkg/comm/render"
	"github.com/bllyanos/weeshlist/internal/pkg/tp"
)

type AuthHandler struct {
	templates *tp.Template
}

func NewAuthHandler(templates *tp.Template) *AuthHandler {
	handler := &AuthHandler{
		templates: templates,
	}
	return handler
}

func (h *AuthHandler) LoginPage(w http.ResponseWriter, r *http.Request) {
	render.Page(h.templates, w, r, "login", nil, nil)
}

// handle login post method
func (h *AuthHandler) LoginAction(w http.ResponseWriter, r *http.Request) {
	// TODO
	fmt.Fprintln(w, "login action")
}

func (h *AuthHandler) RegisterPage(w http.ResponseWriter, r *http.Request) {
	render.Page(h.templates, w, r, "register", nil, nil)
}

// handle register post method
func (h *AuthHandler) RegisterAction(w http.ResponseWriter, r *http.Request) {
	// TODO
	fmt.Fprintln(w, "register action")
}

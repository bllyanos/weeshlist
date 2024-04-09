package multiplexer

import "net/http"

type Mulx struct {
	http.ServeMux
	middlewares []func(http.Handler) http.Handler
}

func NewMulx() *Mulx {
	return &Mulx{middlewares: []func(http.Handler) http.Handler{}}
}

func (m *Mulx) Use(middlewares ...func(http.Handler) http.Handler) {
	m.middlewares = append(m.middlewares, middlewares...)
}

// ServeHTTP implements http.Handler.
func (m *Mulx) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var current http.Handler = &m.ServeMux

	for _, middleware := range m.middlewares {
		current = middleware(current)
	}

	current.ServeHTTP(w, r)
}

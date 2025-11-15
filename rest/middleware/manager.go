package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

type Manager struct {
	middlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		middlewares: []Middleware{},
	}
}

func (m *Manager) Use(mw Middleware) {
	m.middlewares = append(m.middlewares, mw)
}

func (m *Manager) Apply(handler http.Handler, middleware ...Middleware) http.Handler {

	h := handler

	for i := len(middleware) - 1; i >= 0; i-- {
		h = middleware[i](h)
	}

	return h
}

func (m *Manager) ApplyToMux(mux *http.ServeMux) http.Handler {

	h := http.Handler(mux)

	for i := len(m.middlewares) - 1; i >= 0; i-- {
		h = m.middlewares[i](h)
	}

	return h
}

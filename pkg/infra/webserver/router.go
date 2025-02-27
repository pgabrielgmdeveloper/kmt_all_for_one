package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router struct {
	URI            string
	Method         string
	Handler        http.HandlerFunc
	AuthIsRequired bool
}

func NewRouterChi() *chi.Mux {
	return chi.NewRouter()
}

func AddRoutes(r *chi.Mux, baseURI string, routes []Router, middlewares ...func(http.Handler) http.Handler) *chi.Mux {
	if middlewares == nil {

		for _, route := range routes {
			r.Method(route.Method, baseURI+route.URI, route.Handler)
		}

		return r

	}

	for _, route := range routes {

		r.Group(func(r chi.Router) {
			r.Use(middlewares...)
			r.Method(route.Method, baseURI+route.URI, route.Handler)
		})

	}

	return r

}

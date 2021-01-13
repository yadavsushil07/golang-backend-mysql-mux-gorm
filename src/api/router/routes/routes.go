package routes

import (
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url     string
	Method  string
	Handler func(http.ResponseWriter, *http.Request)
}

func Load() []Route {
	routes := userRoutes
	return routes
}

func SetupRoutes(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Url, route.Handler).Methods(route.Method)
	}
	return r
}

func SetupRoutesWithMiddlewares(r *mux.Router) *mux.Router {
	for _, route := range Load() {
		r.HandleFunc(route.Url,
			middlewares.SetMiddlewareLogger(
				middlewares.SetMiddlewareJSON(route.Handler)),
		).Methods(route.Method)
	}
	return r
}

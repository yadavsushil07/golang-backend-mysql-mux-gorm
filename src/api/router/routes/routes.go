package routes

import (
	"api/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Url          string
	Method       string
	Handler      func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func Load() []Route {
	routes := userRoutes
	routes = append(routes, LoginRoutes...)
	routes = append(routes, AdminRoutes...)
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
		if route.AuthRequired {
			r.HandleFunc(route.Url,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(
						middlewares.SetMiddlewareAuthentication(route.Handler))),
			).Methods(route.Method)
		} else {
			r.HandleFunc(route.Url,
				middlewares.SetMiddlewareLogger(
					middlewares.SetMiddlewareJSON(route.Handler)),
			).Methods(route.Method)
		}
	}
	return r
}

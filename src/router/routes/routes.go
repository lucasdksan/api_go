package routes

import (
	"api/src/middlewares"
	"net/http"

	"github.com/gorilla/mux"
)

type Routes struct {
	URI                     string
	Method                  string
	Functionality           func(http.ResponseWriter, *http.Request)
	Requires_authentication bool
}

func Configured(r *mux.Router) *mux.Router {
	routes := user_Routes
	routes = append(routes, route_login)
	routes = append(routes, publications_routes...)

	for _, route := range routes {
		if route.Requires_authentication {
			r.HandleFunc(route.URI, middlewares.Logger(middlewares.Authenticate(route.Functionality))).Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, route.Functionality).Methods(route.Method)
		}
	}

	return r
}

package routes

import (
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

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Functionality).Methods(route.Method)
	}

	return r
}

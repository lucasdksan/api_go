package routes

import (
	"api/src/controllers"
	"net/http"
)

var route_login = Routes{
	URI:                     "/login",
	Method:                  http.MethodPost,
	Functionality:           controllers.Login,
	Requires_authentication: false,
}

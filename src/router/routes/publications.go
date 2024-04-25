package routes

import (
	"api/src/controllers"
	"net/http"
)

var publications_routes = []Routes{
	{
		URI:                     "/publication",
		Method:                  http.MethodPost,
		Functionality:           controllers.Create_Publication,
		Requires_authentication: true,
	},
	{
		URI:                     "/publication",
		Method:                  http.MethodGet,
		Functionality:           controllers.Get_Publication,
		Requires_authentication: true,
	},
	{
		URI:                     "/publication/{id}",
		Method:                  http.MethodPut,
		Functionality:           controllers.Update_Publication,
		Requires_authentication: true,
	},
	{
		URI:                     "/publication/{id}",
		Method:                  http.MethodDelete,
		Functionality:           controllers.Delete_Publication,
		Requires_authentication: true,
	},
}

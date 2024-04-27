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
		URI:                     "/publications",
		Method:                  http.MethodGet,
		Functionality:           controllers.Get_Publications,
		Requires_authentication: true,
	},
	{
		URI:                     "/publication/{id}",
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
	{
		URI:                     "/users/{id}/publications",
		Method:                  http.MethodGet,
		Functionality:           controllers.Search_Posts_From_User,
		Requires_authentication: true,
	},
	{
		URI:                     "/publication/{id}/like",
		Method:                  http.MethodPost,
		Functionality:           controllers.Like_Post,
		Requires_authentication: true,
	},
	{
		URI:                     "/publication/{id}/dislike",
		Method:                  http.MethodPost,
		Functionality:           controllers.Dislike_Post,
		Requires_authentication: true,
	},
}

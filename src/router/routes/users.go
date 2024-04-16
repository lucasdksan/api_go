package routes

import (
	"api/src/controllers"
	"net/http"
)

var user_Routes = []Routes{
	{
		URI:                     "/users",
		Method:                  http.MethodPost,
		Functionality:           controllers.Create_User,
		Requires_authentication: false,
	},
	{
		URI:                     "/users",
		Method:                  http.MethodGet,
		Functionality:           controllers.Get_Users,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodGet,
		Functionality:           controllers.Get_User,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodPut,
		Functionality:           controllers.Update_User,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}",
		Method:                  http.MethodDelete,
		Functionality:           controllers.Delete_User,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}/follow",
		Method:                  http.MethodPost,
		Functionality:           controllers.Follow_user,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}/unfollow",
		Method:                  http.MethodPost,
		Functionality:           controllers.Un_follow_user,
		Requires_authentication: true,
	},
	{
		URI:                     "/users/{id}/followers",
		Method:                  http.MethodGet,
		Functionality:           controllers.Get_Followers,
		Requires_authentication: true,
	},
}

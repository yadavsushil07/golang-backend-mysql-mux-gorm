package routes

import (
	"api/controllers"
	"net/http"
)

var userRoutes = []Route{
	Route{
		Url:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Url:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
	Route{
		Url:     "/user",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		Url:     "/user/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		Url:     "/user/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
}

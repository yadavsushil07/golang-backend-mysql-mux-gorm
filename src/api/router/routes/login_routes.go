package routes

import (
	"api/controllers"
	"net/http"
)

var LoginRoutes = []Route{
	Route{
		Url:          "/login",
		Method:       http.MethodPost,
		Handler:      controllers.Login,
		AuthRequired: false,
	},
	Route{
		Url:          "/adminlogin",
		Method:       http.MethodPost,
		Handler:      controllers.AdminLogin,
		AuthRequired: false,
	},
	Route{
		Url:          "/registration",
		Method:       http.MethodPost,
		Handler:      controllers.Registration,
		AuthRequired: false,
	},
}

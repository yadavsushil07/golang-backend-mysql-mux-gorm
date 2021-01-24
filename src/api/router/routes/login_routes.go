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
	Route{
		Url:          "/forgotpassword",
		Method:       http.MethodPost,
		Handler:      controllers.ForgotPassword,
		AuthRequired: false,
	},
	Route{
		Url:          "/newpassword",
		Method:       http.MethodPost,
		Handler:      controllers.NewPassword,
		AuthRequired: false,
	},
}

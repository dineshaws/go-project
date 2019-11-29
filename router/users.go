package router

import (
	"github.com/gorilla/mux"

	"github.com/dineshaws/go-project/user"
)

var userDao = user.DAO{}

var userController user.Controller

func init() {
	cnfg.Read()
	userDao.Server = cnfg.Server
	userDao.Database = cnfg.Database
	userDao.Connect()
	userController = user.Controller{userDao}
}

var userRoutes = Routes{
	Route{
		"AddUser",
		"POST",
		"/api/v1/users",
		userController.AddUser,
	},
	Route{
		"Login",
		"POST",
		"/api/v1/users/login",
		userController.LoginUser,
	},
	Route{
		"Profile",
		"GET",
		"/api/v1/users/{id}",
		userController.GetUser,
	},
}

// SetUsersRoutes function to set users routes
func SetUsersRoutes(r *mux.Router) *mux.Router {
	for _, route := range userRoutes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(route.HandleFunc)
	}

	return r
}

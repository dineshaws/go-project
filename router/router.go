package router

import (
	"net/http"

	"github.com/dineshaws/go-project/config"
	"github.com/gorilla/mux"
)

// Route map
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

// Routes List of map
type Routes []Route

var cnfg = config.Config{}

// InitRouter function
func InitRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r = SetAlbumRouters(r)
	r = SetUsersRoutes(r)
	return r
}

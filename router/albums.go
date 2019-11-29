package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/dineshaws/go-project/album"
	"github.com/dineshaws/go-project/config"
)

// Route map
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

var cnfg = config.Config{}

var albumDao = album.AlbumDAO{}

var controller album.Controller

// init function
func init() {
	fmt.Println("Router package init")
	cnfg.Read()
	albumDao.Server = cnfg.Server
	albumDao.Database = cnfg.Database
	albumDao.Connect()
	controller = album.Controller{albumDao}

}

// Routes List of map
type Routes []Route

var routes = Routes{
	Route{
		"GetAllAlbums",
		"GET",
		"/api/v1/albums",
		controller.GetAllAlbum,
	},
	Route{
		"AddAlbum",
		"POST",
		"/api/v1/albums",
		controller.AddAlbum,
	},
	Route{
		"UpdateAlbum",
		"PUT",
		"/api/v1/albums/{id}",
		controller.UpdateAlbum,
	},
	Route{
		"GetAlbum",
		"GET",
		"/api/v1/albums/{id}",
		controller.GetAlbum,
	},
	Route{
		"DeleteAlbum",
		"DELETE",
		"/api/v1/albums/{id}",
		controller.DeleteAlbum,
	},
}

// SetAlbumRouters function to create album routes
func SetAlbumRouters(r *mux.Router) *mux.Router {

	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)

	}
	return r

}

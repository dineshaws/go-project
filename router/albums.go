package router

import (
	"fmt"

	"github.com/gorilla/mux"

	"github.com/dineshaws/go-project/album"
)

var albumDao = album.AlbumDAO{}

var controller album.Controller

// init function
func init() {
	fmt.Println("Album Router init")
	cnfg.Read()
	albumDao.Server = cnfg.Server
	albumDao.Database = cnfg.Database
	albumDao.Connect()
	controller = album.Controller{albumDao}

}

var albumRoutes = Routes{
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

	for _, route := range albumRoutes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)

	}
	return r

}

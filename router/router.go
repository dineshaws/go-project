package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/dineshaws/go-project/album"
	. "github.com/dineshaws/go-project/config"
)

// Route map
type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

var config = Config{}

var albumDao = AlbumDAO{}

var controller Controller

// init function
func init() {
	fmt.Println("Router package init")
	config.Read()
	albumDao.Server = config.Server
	albumDao.Database = config.Database
	albumDao.Connect()
	controller = Controller{albumDao}

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
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "not yet")
		},
	},
	Route{
		"UpdateAlbum",
		"PUT",
		"/api/v1/albums",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "not yet")
		},
	},
	Route{
		"GetAlbum",
		"GET",
		"/api/v1/albums",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "not yet")
		},
	},
	Route{
		"DeleteAlbum",
		"DELETE",
		"/api/v1/albums",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintln(w, "not yet")
		},
	},
}

// NewRouter function
func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		r.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(route.HandleFunc)

	}
	return r

}

package router

import (
	"github.com/gorilla/mux"
)

// InitRouter function
func InitRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	r = SetAlbumRouters(r)
	return r
}

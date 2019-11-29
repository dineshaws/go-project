package album

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2"

	"github.com/dineshaws/go-project/common"
)

// Controller map
type Controller struct {
	AlbumDAO AlbumDAO
}

// GetAllAlbum controller method
func (c *Controller) GetAllAlbum(w http.ResponseWriter, r *http.Request) {
	albums, err := c.AlbumDAO.GetAll()
	if err != nil {
		common.RespondWithError(w, http.StatusInternalServerError, "An unexpected error has occurred", err)
		return
	}
	common.RespondWithJSON(w, http.StatusOK, albums)

}

// GetAlbum controller method
func (c *Controller) GetAlbum(w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	message := "An unexpected error has occurred"

	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		code = http.StatusBadRequest
		message = "Wrong object id"
		common.RespondWithError(w, code, message, errors.New("Object id Incorrect"))
		return
	}
	album, err := c.AlbumDAO.GetByID(id)
	if err != nil {
		if err == mgo.ErrNotFound {
			code = http.StatusNotFound
			message = "No record exist"
		}
		common.RespondWithError(w, code, message, err)
		return
	}
	common.RespondWithJSON(w, http.StatusOK, album)
}

// UpdateAlbum controller method
func (c *Controller) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	message := "An unexpected error has occurred"

	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		code = http.StatusBadRequest
		message = "Wrong object id"
		common.RespondWithError(w, code, message, errors.New("Object id Incorrect"))
		return
	}

	defer r.Body.Close()
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read content
	if err != nil {
		common.RespondWithError(w, code, message, err)
		return
	}

	var album Album
	if err := json.Unmarshal(body, &album); err != nil {
		message = "Error while unmarshalling request"
		common.RespondWithError(w, code, message, err)
		return
	}

	album.ID = bson.ObjectIdHex(id)
	info, err := c.AlbumDAO.UpdateByID(album)
	fmt.Println(info)
	if err != nil {
		message = "Error while updating doc"
		common.RespondWithError(w, code, message, err)
		return
	}
	common.RespondWithJSON(w, http.StatusOK, album)
}

// DeleteAlbum controller method
func (c *Controller) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	message := "An unexpected error has occurred"

	vars := mux.Vars(r)
	id := vars["id"]
	if !bson.IsObjectIdHex(id) {
		code = http.StatusBadRequest
		message = "Wrong object id"
		common.RespondWithError(w, code, message, errors.New("Object id Incorrect"))
		return
	}

	if err := c.AlbumDAO.DeleteByID(bson.ObjectIdHex(id)); err != nil {
		message = "Error while updating doc"
		common.RespondWithError(w, code, message, err)
		return
	}
	common.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "success"})
}

// AddAlbum controller method
func (c *Controller) AddAlbum(w http.ResponseWriter, r *http.Request) {
	code := http.StatusInternalServerError
	message := "An unexpected error has occurred"

	defer r.Body.Close()
	var album Album
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		message = "Error while reading request"
		common.RespondWithError(w, code, message, err)
		return
	}
	if err := json.Unmarshal(body, &album); err != nil {
		message = "Error while unmarshalling request"
		common.RespondWithError(w, code, message, err)
		return
	}
	album.ID = bson.NewObjectId() // add _id before insertion
	if err := c.AlbumDAO.Add(album); err != nil {
		message = "Error while inserting doc"
		common.RespondWithError(w, code, message, err)
		return
	}
	common.RespondWithJSON(w, http.StatusOK, album)

}

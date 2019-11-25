package album

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Controller map
type Controller struct {
	AlbumDAO AlbumDAO
}

// GetAllAlbum controller method
func (c *Controller) GetAllAlbum(w http.ResponseWriter, r *http.Request) {
	albums, err := c.AlbumDAO.GetAll()
	if err != nil {
		fmt.Fprintln(w, err)
	}
	respondWithJSON(w, http.StatusOK, albums)

}

// respondWithJson function
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

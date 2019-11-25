package album

import (
	"log"

	"gopkg.in/mgo.v2"
)

// AlbumDAO map
type AlbumDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Connect Function
func (dao *AlbumDAO) Connect() {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dao.Server)
}

// GetAll albums
func (dao *AlbumDAO) GetAll() ([]string, interface{}) {
	var albums = []string{"dinesh"}
	return albums, nil
}

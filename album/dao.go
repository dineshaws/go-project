package album

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// AlbumDAO map
type AlbumDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION name is albums
var COLLECTION = "albums"

// Connect Function
func (dao *AlbumDAO) Connect() {
	// var logger = log.New(os.Stdout, "logger: ", log.Ldate|log.Ltime|log.Lmicroseconds)
	// mgo.SetLogger(logger)
	// mgo.SetDebug(true)
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dao.Database)
}

// GetAll albums
func (dao *AlbumDAO) GetAll() (Albums, error) {
	var albums = []Album{}
	err := db.C(COLLECTION).Find(bson.M{}).All(&albums)
	return albums, err
}

// GetByID get single document by id
func (dao *AlbumDAO) GetByID(id string) (Album, error) {
	var album Album
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&album)
	return album, err
}

// Add function to create a  new document in collection
func (dao *AlbumDAO) Add(album Album) error {
	err := db.C(COLLECTION).Insert(album)
	return err
}

// UpdateByID updated the doc
func (dao *AlbumDAO) UpdateByID(album Album) (*mgo.ChangeInfo, error) {
	info, err := db.C(COLLECTION).UpsertId(album.ID, album)
	return info, err
}

// DeleteByID delete document from collection
func (dao *AlbumDAO) DeleteByID(id bson.ObjectId) error {
	err := db.C(COLLECTION).RemoveId(id)
	return err
}

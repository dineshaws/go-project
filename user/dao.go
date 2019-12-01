package user

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// DAO struct
type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// COLLECTION name is users
const COLLECTION = "users"

// Connect method to connect db
func (dao *DAO) Connect() {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dao.Database)
}

// Add method to create new user
func (dao *DAO) Add(user User) error {
	err := db.C(COLLECTION).Insert(user)
	return err
}

// FindByUsername method to check username existance
func (dao *DAO) FindByUsername(username string) (int, error) {
	count, err := db.C(COLLECTION).Find(bson.M{"username": username}).Count()
	return count, err
}

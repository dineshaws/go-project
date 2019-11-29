package user

import (
	"log"

	"gopkg.in/mgo.v2"
)

// DAO struct
type DAO struct {
	Server   string
	Database string
}

var db *mgo.Database

// Connect method to connect db
func (dao *DAO) Connect() {
	session, err := mgo.Dial(dao.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(dao.Database)
}

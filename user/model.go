package user

import (
	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Username  string        `bson:"username" json:"username"`
	FirstName string        `bson:"firstname" json:"firstname"`
	LastName  string        `bson:"lastname" json:"lastname"`
	Password  string        `bson:"password" json:"password,omitempty"`
	Token     string        `bson:"token" json:"token"`
}

// Users type declared as array
type Users []User

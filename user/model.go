package user

import (
	"gopkg.in/mgo.v2/bson"
)

// User model
type User struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Username  string        `bson:"username" json:"username"`
	FirstName string        `json:"firstname" json:"firstname"`
	LastName  string        `json:"lastname" json:"lastname"`
	Password  string        `json:"password" json:"password"`
	Token     string        `json:"token" json:"token"`
}

// Users type declared as array
type Users []User

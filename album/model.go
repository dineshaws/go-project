package album

import (
	"gopkg.in/mgo.v2/bson"
)

// Album represents a music album
type Album struct {
	ID     bson.ObjectId `bson:"_id" json:"id,omitempty"`
	Title  string        `bson:"title" json:"title"`
	Artist string        `bson:"artist" json:"artist"`
	Year   int           `bson:"year" json:"year"`
}

// Albums is an array of Album
type Albums []Album

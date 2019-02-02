package song

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Song Struct
type Song struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	User      string        `bson:"user" json:"user"`
	Preferred bool          `bson:"preferred" json:"preferred"`
	Name      string        `bson:"name" json:"name"`
	Source    string        `bson:"source" json:"source"`
	Link      string        `bson:"link" json:"link"`
	Timestamp time.Time     `bson:"timestamp" json:"timestamp"`
}

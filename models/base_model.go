package models

import (
	"time"

	"labix.org/v2/mgo/bson"
)

type BaseModel struct {
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
}
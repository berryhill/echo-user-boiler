package models

import (
	"time"

	"github.com/user-boiler/store"

	"labix.org/v2/mgo/bson"
)

type Printer struct {
	Id 		bson.ObjectId 		`json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time		`json:"time",bson:"time,omitempty"`
	Technology 	string			`json:"technology",bson:"technology,omitempty"`
	MaxX		float32			`json:"max_x",bson:"max_x,omitempty"`
	MaxY		float32			`json:"max_y",bson:"max_y,omitempty"`
	MaxZ		float32			`json:"max_z",bson:"max_z,omitempty"`
	Colors 		[]string		`json:"colors",bson:"colors,omitempty"`
	Resolutions	[]int			`json:"resolutions",bson:"resolutions,omitempty"`
}

func NewPrinter() *Printer {
	po := new(Printer)
	po.Id = bson.NewObjectId()
	po.Timestamp = time.Now()

	return po
}

func (po *Printer) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "printers")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(po)
	if err != nil {
		return err
	}

	return nil
}

func (po *Printer) Update() error {
	//TODO implement; this is where the actual db update should be

	return nil
}

func (po *Printer) AddColor(color string) error {
	//TODO implement; need to be able to update the PrintOptions

	return nil
}

func (po *Printer) AddResolution(resolution int) error {
	//TODO implement; need to add res

	return nil
}

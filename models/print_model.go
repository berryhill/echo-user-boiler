package models

import (
	"time"
	"errors"

	"github.com/user-boiler/store"

	"labix.org/v2/mgo/bson"
	"encoding/binary"
)

//TODO implement; would like to use an enum for the Status instead of a string
//type Enum string
//const (
//	NotCompleted Enum = ""
//	Completed
//)

type Print struct {
	Id 		bson.ObjectId 		`json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time		`json:"time",bson:"time,omitempty"`
	CustomerId 	bson.ObjectId		`json:"customer_id",bson:"customer_id,omitempty"`
	Status 		string			`json:"status",bson:"status,omitempty"`
	Stl 		binary.ByteOrder	`json:"stl",bson:"stl,omitempty"`
	Resolution	int			`json:"resolution",bson:"resolution,omitempty"`
	Volume 		float32			`json:"volume",bson:"volume,omitempty"`
	Infill		float32			`json:"infill",bson:"infill,omitempty"`
	Price		float32			`json:"price",bson:"price,omitempty"`
}

func NewPrint() *Print {
	p := new(Print)
	p.Id = bson.NewObjectId()
	p.Timestamp = time.Now()
	p.Status = "New Order"
	//p.Stl = stl

	return p
}

func (p *Print) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "prints")
	if err != nil {
		panic(err)
	}

	print := Print {
		Id:		p.Id,
		Timestamp:	p.Timestamp,
	}
	err = collection.Insert(print)
	if err != nil {
		return err
	}

	return nil
}

func (p *Print) ChangeStatus(status string) error {
	//TODO implement an enum to check it status's exist before changing
	p.Status = status

	return nil
}

func FindPrintById(id string) (*Print, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "prints")
	if err != nil {
		//panic(err)
		return &Print{}, err
	}

	print := Print{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&print)
	if err != nil {
		//panic(err)
		return &print, err
	}

	return &print, err
}

func GetAllPrints() ([]*Print, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "prints")
	if err != nil {
		panic(err)
	}

	prints := []*Print{}
	err = collection.Find(nil).All(&prints)

	return prints, err
}

func (p *Print) checkResolution(resolution int) error {
	if resolution < 100 || resolution > 400 {
		return errors.New("Unsupported Resolution")
	} else {
		if resolution % 10 != 0 {
			return errors.New("Unsupported Resolution")
		} else {
			return nil
		}
	}
}

func (p *Print) calculateVolume() (float32, error) {
	//TODO implement

	return 4.20, nil
}

func (p *Print) calculatePrice() (float32, error) {
	//TODO implement

	return 10.00, nil
}
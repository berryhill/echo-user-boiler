package models

import (
	"time"
	"log"

	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
)

type User struct {
	//BaseModel
	Id 		bson.ObjectId          `json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	       `json:"time",bson:"time,omitempty"`
	Username	string           `json:"username",bson:"username,omitempty"`
	Password	string           `json:"password",bson:"password,omitempty"`
}

func NewUser(username string, password string) *User {
	u := new(User)
	u.Id = bson.NewObjectId()
	u.Username = username
	u.Password = password
	return u
}

func (u *User) Save() error {
	session, err := mgo.Dial("localhost")
	if err != nil {
		return err
	}
	defer session.Close()

	c := session.DB("test").C("users")
	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	err = c.Insert(&User{Id: u.Id,
		Timestamp: u.Timestamp,
		Username: u.Username,
		Password: u.Password})
	if err != nil {
		panic(err)
	}

	return nil
}

func FindUser(username string) (*User, error) {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()

	c := session.DB("test").C("users")
	index := mgo.Index{
		Key:        []string{"username"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}

	user := User{}
	err = c.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		log.Fatal(err)
	}

	if user.Id == "" {
		err := mgo.ErrNotFound
		return &user, err
	}

	return &user, err
}
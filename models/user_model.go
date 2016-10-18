package models

import (
	"time"

	"labix.org/v2/mgo/bson"
	"github.com/user-boiler/store"
)

type User struct {
	//BaseModel
	Id 		bson.ObjectId 	`json:"id",bson:"_id,omitempty"`
	Timestamp 	time.Time	`json:"time",bson:"time,omitempty"`
	Username	string		`json:"username",bson:"username,omitempty"`
	Password	string		`json:"password",bson:"password,omitempty"`
}

func NewUser(username string, password string) *User {
	u := new(User)
	u.Id = bson.NewObjectId()
	u.Timestamp = time.Now()
	u.Username = username
	u.Password = password

	return u
}

func (u *User) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		panic(err)
	}

	user := User {
		Id:		u.Id,
		Timestamp:	u.Timestamp,
		Username:	u.Username,
		Password: 	u.Password,
	}

	err = collection.Insert(user)
	if err != nil {
		return err
	}

	return nil
}

func FindUserByName(username string) (*User, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		panic(err)
	}

	user := User{}
	err = collection.Find(bson.M{"username": username}).One(&user)
	if err != nil {
		return &user, err
	}

	return &user, err
}

func FindUserById(id string) (*User, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		//panic(err)
		return &User{}, err
	}

	user := User{}
	err = collection.Find(bson.M{"id": bson.ObjectIdHex(id)}).One(&user)
	if err != nil {
		//panic(err)
		return &user, err
	}

	return &user, err
}

func GetAllUsers() ([]*User, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		panic(err)
	}

	users := []*User{}
	err = collection.Find(nil).All(&users)

	return users, err
}
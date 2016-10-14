package models

import (
	"encoding/json"

	"gopkg.in/mgo.v2/bson"
	"github.com/user-boiler/store"
)

type Question struct {
	Id		bson.ObjectId 	`json:"id",bson:"_id,omitempty"`
	Text 		string		`json:"text",bson:"text,omitempty"`
	//AnswerId	*bson.ObjectId	`json:"id",bson:"_id,omitempty"`
}

func NewQuestion(text string) *Question {
	q := new(Question)
	q.Id = bson.NewObjectId()
	q.Text = text

	return q
}

func (q *Question) MarshalJson() ([]byte, error) {
	return json.Marshal(q)
}

func (q *Question) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "questions")
	if err != nil {
		panic(err)
	}

	err = collection.Insert(&Question{
		Id: q.Id,
		Text: q.Text})
	if err != nil {
		panic(err)
	}

	return nil
}

func FindQuestion(id string) (Question, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "questions")
	if err != nil {
		panic(err)
	}

	question := Question{}
	err = collection.Find(bson.M{"_id": id}).One(&question)
	if err != nil {
		return question, err
	}

	return question, err
}

func Questions() ([]Question, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "questions")
	if err != nil {
		panic(err)
	}

	questions := []Question{}
	err = collection.Find(nil).All(&questions)
	if err != nil {
		return questions, err
	}

	return questions, err
}

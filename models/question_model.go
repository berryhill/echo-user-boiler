package models

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/user-boiler/store"
)

type Question struct {
	Id		bson.ObjectId
	Text 		string
	//AnswerId	*bson.ObjectId
}

func NewQuestion(text string) *Question {
	q := new(Question)
	q.Id = bson.NewObjectId()
	q.Text = text

	return q
}

func (q *Question) Save() error {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}
	collection, err := store.ConnectToCollection(session, "checkbox_questions")
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

func (q *Question) GetId() *bson.ObjectId {
	return &q.Id
}

func (q *Question) GetText() string {
	return q.Text
}

//func (cbq *CheckboxQuestion) GetAnswer() string {
//	return "Implement"
//}

func FindQuestion(id string) (Question, error) {
	session, err := store.ConnectToDb()
	defer session.Close()
	if err != nil {
		panic(err)
	}

	collection, err := store.ConnectToCollection(session, "users")
	if err != nil {
		panic(err)
	}

	question := Question{}
	err = collection.Find(bson.M{"id": id}).One(&question)
	if err != nil {
		return question, err
	}

	return question, err
}

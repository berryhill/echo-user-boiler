package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
	"fmt"
)

type Survey struct {
	Id 		bson.ObjectId
	Timestamp	time.Time
	SurveyorLink	string
	QuestionIds	[]*bson.ObjectId
}

func NewSurvey() *Survey {
	s := new(Survey)
	s.Id = bson.NewObjectId()
	s.Timestamp = time.Now()

	return s
}

func (s *Survey) AddQuestion(q *Question) error {
	s.QuestionIds = append(s.QuestionIds, q.GetId())

	return nil
}

func (s *Survey) GetQuestionWithId(q_id string) (*Question, error) {
	question := new(Question)

	return question, nil
}

func (s *Survey) Print() {
	fmt.Println("Survey ID %s", s.Id)
	for k:=0; k<len(s.QuestionIds); k++ {
		fmt.Println("Question: ")
		//question, _ := s.GetQuestionWithId(s.QuestionIds[k])
		//fmt.Println(question.GetQuestionText())
	}
}

func InitTestSurvey() *Survey {
	survey := NewSurvey()
	question := NewQuestion("This is a Question")
	survey.AddQuestion(question)

	return survey
}


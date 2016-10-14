package server

import (
	"net/http"

	"github.com/user-boiler/models"

	"github.com/labstack/echo"
)

func CreateQuestion(c echo.Context) error {
	text := c.FormValue("text")

	question := models.NewQuestion(text)
	err := question.Save()
	if err != nil {
		return err
	}

	return nil
}

func GetQuestion(c echo.Context) error {
	id := c.Param("id")

	question, err := models.FindQuestion(id)
	if err != nil {
		panic(err)
	}

	if question.Id != "" /*&& user.Username != "" */ {
		return c.JSON(http.StatusOK, question)
	} else {
		return c.JSON(http.StatusNotFound, "not found")
	}
}

func GetAllQuestions(c echo.Context) error {
	questions, err := models.Questions()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, questions)
}

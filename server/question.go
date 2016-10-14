package server

import (
	"github.com/user-boiler/models"

	"github.com/labstack/echo"
	"net/http"
)

func CreateQuestion(c echo.Context) error {
	qtext := c.FormValue("qtext")

	question := models.NewQuestion(qtext)
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

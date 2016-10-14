package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
	"github.com/user-boiler/models"
)

func Run() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	//e.Use(middleware.Static("/static"))

	e.Static("/", "static")

	//e.GET("/", accessible)
	r.GET("", restricted)

	e.GET("/user/:username", GetUser)
	e.GET("/question/:id", GetQuestion)
	e.GET("/questions", GetAllQuestions)

	e.POST("/user", CreateUser)
	e.POST("/login", Login)
	e.POST("/question", CreateQuestion)

	survey := models.InitTestSurvey()
	survey.Print()
	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
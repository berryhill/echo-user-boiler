package server

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/engine/standard"
)

func Run() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Restricted group
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	e.POST("/user", CreateUser)
	e.GET("/user", GetUser)
	e.POST("/login", Login)
	e.GET("/", accessible)
	r.GET("", restricted)

	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
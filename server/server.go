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
	r := e.Group("/restricted")
	r.Use(middleware.JWT([]byte("secret")))

	e.GET("/", accessible)
	r.GET("", restricted)

	e.GET("/user/:id", GetUserById)
	e.GET("/user/name/:username", GetUserByUsername)
	e.POST("/user", CreateUser)
	e.PUT("/user/:id", UpdateUser)
	e.GET("/users", GetAllUsers)
	e.POST("/login", Login)
	e.DELETE("/user/:id", DeleteUser)


	fmt.Println("Server now running on port: 1323")
	e.Run(standard.New(":1323"))
}
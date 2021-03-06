package server

import (
	"time"
	"net/http"
	"fmt"

	"github.com/midi-survey/models"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
)

func CreateUser(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user := models.NewUser(username, password)
	err := user.Save()
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func Login(c echo.Context) error {
	username := c.FormValue("username")
	password := c.FormValue("password")

	user, err := models.FindUser(username)
	if err != nil {
		panic(err)
	}

	if user.Password == password {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = "Jon Snow"
		claims["admin"] = true
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte("secret"))
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}

	return echo.ErrUnauthorized
}

func GetUser(c echo.Context) error {
	username := c.FormValue("username")
	//password := c.FormValue("password")

	user, err := models.FindUser(username)
	if err != nil {
		panic(err)
	}

	if user.Password != "" {
		return c.JSON(http.StatusOK, user)
	} else {
		return c.JSON(http.StatusNotFound, "not found")
	}
}


package server

import (
	"time"
	"net/http"
	"encoding/json"
	"fmt"

	"github.com/user-boiler/models"
	"io/ioutil"

	"github.com/labstack/echo"
	"github.com/dgrijalva/jwt-go"
	"labix.org/v2/mgo/bson"
	log "github.com/cihub/seelog"
)

func CreateUser(c echo.Context) error {
	method := c.Request().Method()
	uri := c.Request().URI()
	log.Debugf("%s %s", method, uri)

	json_body, err := ioutil.ReadAll(c.Request().Body())
	user := models.User{}
	err = json.Unmarshal(json_body, &user)
	if err != nil {
		fmt.Println(err)
	}

	user.Timestamp = time.Now()
	user.Id = bson.NewObjectId()

	err = user.Save()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, user)
}

func GetUser(c echo.Context) error {
	username := c.Param("username")

	user, err := models.FindUserByName(username)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, user)
}

func GetUserById(c echo.Context) error {
	id := c.Param("id")

	user, err := models.FindUserById(id)
	if err != nil {
		panic(err)

	}

	return c.JSON(http.StatusOK, user)
}

func GetAllUsers(c echo.Context) error {
	users, err := models.GetAllUsers()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, users)
}

func UpdateUser(c echo.Context) error {
	id := c.Param("id")
	fmt.Println(id)

	//TODO implement

	return nil
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	user, err := models.FindUserById(id)
	if err != nil {
		panic(err)

	}

	err = user.Delete()
	if err != nil {
		panic(err)

	}

	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	method := c.Request().Method()
	uri := c.Request().URI()
	log.Debugf("%s %s", method, uri)

	json_body, err := ioutil.ReadAll(c.Request().Body())
	u := models.User{}
	err = json.Unmarshal(json_body, &u)
	if err != nil {
		fmt.Println(err)
	}

	user, err := models.FindUserByName(u.Username)
	if err != nil {
		panic(err)
	}

	if u.Password == user.Password {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["name"] = user.Username
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


package server

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/user-boiler/models"

	"github.com/labstack/echo"
	log "github.com/cihub/seelog"
)

func CreatePrint(c echo.Context) error {
	method := c.Request().Method()
	uri := c.Request().URI()
	log.Debugf("%s %s", method, uri)

	json_body, err := ioutil.ReadAll(c.Request().Body())
	p := models.Print{}
	err = json.Unmarshal(json_body, &p)
	if err != nil {
		fmt.Println(err)
	}

	print := models.NewPrint()

	//TODO implement; will need to combine parsed json with NewPrint

	err = print.Save()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, print)
}

func GetPrintById(c echo.Context) error {
	id := c.Param("id")
	print, err := models.FindPrintById(id)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, print)
}

func GetAllPrints(c echo.Context) error {
	prints, err := models.GetAllPrints()
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, prints)
}
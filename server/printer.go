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

func CreatePrinter(c echo.Context) error {
	method := c.Request().Method()
	uri := c.Request().URI()
	log.Debugf("%s %s", method, uri)

	json_body, err := ioutil.ReadAll(c.Request().Body())
	p := models.Printer{}
	err = json.Unmarshal(json_body, &p)
	if err != nil {
		fmt.Println(err)
	}

	printer := models.NewPrinter()
	//TODO implement; will need to combine parsed json with NewPrinter

	err = printer.Save()
	if err != nil {
		fmt.Println(err)
	}

	return c.JSON(http.StatusOK, printer)
}

//func GetPrintertById(c echo.Context) error {
//	id := c.Param("id")
//	print, err := models.FindPrinterById(id)
//	if err != nil {
//		panic(err)
//	}
//
//	return c.JSON(http.StatusOK, print)
//}

//func GetAllPrinters(c echo.Context) error {
//	printers, err := models.GetAllPrinters()
//	if err != nil {
//		panic(err)
//	}
//
//	return c.JSON(http.StatusOK, printers)
//}

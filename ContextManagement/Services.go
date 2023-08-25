package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Services struct {
}

func Init() *Services {
	return &Services{}
}

func (service *Services) HandleGetRequest(c echo.Context) error {
	return c.String(http.StatusOK, "Server is up")
}

func (service *Services) HandlePostRequest(c echo.Context) error {
	fmt.Println(">> ", c.QueryParam("message"))
	return c.String(http.StatusOK, "")
}

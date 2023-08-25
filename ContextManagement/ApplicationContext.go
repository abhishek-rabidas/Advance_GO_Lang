package main

import (
	"github.com/labstack/echo/v4"
)

type ApplicationContext struct {
	Server   *echo.Echo
	Services *Services
}

func CreateApplicationContext(services *Services) *ApplicationContext {
	e := echo.New()
	ctx := ApplicationContext{
		Server:   e,
		Services: services,
	}

	routes(&ctx)

	err := e.Start(":5050")

	if err != nil {
		panic("Server not running")
	}

	return &ctx
}

func routes(ctx *ApplicationContext) {
	e := ctx.Server

	e.GET("/", ctx.Services.HandleGetRequest)

	e.POST("/", ctx.Services.HandlePostRequest)
}

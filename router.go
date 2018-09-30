package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func createRouter() *echo.Echo {
	e := echo.New()
	applyMiddlewares(e)
	applyRoutes(e)
	return e
}

func applyRoutes(e *echo.Echo) {
	e.File("/", "client/index.html")
}

func applyMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
}

package main

import "github.com/labstack/echo"

func createRouter() *echo.Echo {
	e := echo.New()
	applyRoutes(e)
	return e
}

func applyRoutes(e *echo.Echo) {
	e.File("/", "client/index.html")
}

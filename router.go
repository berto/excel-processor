package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/berto/excel-processor/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/tealeg/xlsx"
)

func createRouter() *echo.Echo {
	e := echo.New()
	applyMiddlewares(e)
	applyRoutes(e)
	return e
}

func applyRoutes(e *echo.Echo) {
	e.File("/", "client/index.html")
	e.GET("/migrate", migrateDB)
	e.GET("/seed", seedShips)
	e.GET("/budget/:code", getBudget)
	e.POST("/ship", parseShipData)
}

func applyMiddlewares(e *echo.Echo) {
	e.Use(middleware.Logger())
}

func migrateDB(c echo.Context) error {
	message := "Migration Successful"
	err := db.RunMigrations()
	if err != nil {
		message = err.Error()
	}
	return c.String(http.StatusOK, message)
}

func seedShips(c echo.Context) error {
	message := "Seeding Successful"
	err := db.SeedShips()
	if err != nil {
		message = err.Error()
	}
	return c.String(http.StatusOK, message)
}

func getBudget(c echo.Context) error {
	code := c.Param("code")
	budgetDataResponse := db.GetBudgetDataFromCode(code)
	return c.JSON(http.StatusOK, budgetDataResponse)
}

func parseShipData(c echo.Context) error {
	file, err := c.FormFile("ship")
	if err != nil {
		return err
	}
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	data, err := ioutil.ReadAll(src)
	if err != nil {
		log.Fatal(err)
	}

	xlFile, err := xlsx.OpenBinary(data)
	if err != nil {
		return err
	}

	code, err := parseExcelData(xlFile, file.Filename)
	if err != nil {
		return err
	}
	return c.Redirect(http.StatusMovedPermanently, "/budget/"+code)
}

package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := createRouter()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Error(err)
	}

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

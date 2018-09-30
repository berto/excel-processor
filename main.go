package main

import (
	"os"

	"github.com/joho/godotenv"
)

func main() {
	e := createRouter()

	err := godotenv.Load()
	if err != nil {
		e.Logger.Fatal(err)
	}

	port := os.Getenv("PORT")
	e.Logger.Fatal(e.Start(":" + port))
}

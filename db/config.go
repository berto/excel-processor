package main

import (
	"os"

	"github.com/joho/godotenv"
)

func generateDatabaseURL() (string, error) {
	err := godotenv.Load()
	if err != nil {
		return "", err
	}
	DB_HOST := os.Getenv("DB_HOST")
	DB_USERNAME := os.Getenv("DB_USERNAME")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	return DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":3306)" + "/" + DB_NAME, nil
}
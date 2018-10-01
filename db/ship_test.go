package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestShip(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Errorf("Failed to load env with error %d.", err)
	}

	DB_NAME := os.Getenv("DB_TEST_NAME")
	os.Setenv("DB_NAME", DB_NAME)

	ship := Ship{
		ID:   1,
		Name: "Titanic",
		Year: 1900,
	}

	err = AddShip(ship)
	if err != nil {
		t.Errorf("Failed to insert ship with error %d.", err)
	}
}

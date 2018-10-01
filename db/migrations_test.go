package db

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestMigrations(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Errorf("Failed to load env with error %d.", err)
	}

	DB_NAME := os.Getenv("DB_TEST_NAME")
	os.Setenv("DB_NAME", DB_NAME)

	err = RunMigrations()
	if err != nil {
		t.Errorf("Failed to run migrations with error %d.", err)
	}
}

package main

import (
	_ "github.com/go-sql-driver/mysql"

	"github.com/jmoiron/sqlx"
)

func RunMigrations() error {
	dbURL, err := generateDatabaseURL()
	if err != nil {
		return err
	}

	db, err := sqlx.Connect("mysql", dbURL)
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.MustBegin()

	tx.MustExec(createShipTable)

	tx.MustExec(`TRUNCATE TABLE ship;`)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

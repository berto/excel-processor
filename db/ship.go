package main

import "github.com/jmoiron/sqlx"

type Ship struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Size int    `db:"size"`
	Year int    `db:"year"`
}

func insertShip(ship Ship) error {
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

	tx.NamedExec(insertShipQuery, ship)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

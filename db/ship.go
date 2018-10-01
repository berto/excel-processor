package db

type Ship struct {
	ID   int    `db:"id"`
	Name string `db:"name"`
	Size int    `db:"size"`
	Year int    `db:"year"`
}

func AddShip(ship Ship) error {
	db, err := getDB()
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

func FindShipByName(name string) (ship Ship, err error) {
	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Get(&ship, getShipByNameQuery, name)
	if err != nil {
		return
	}

	return
}

func FindShipByID(id int) (ship Ship, err error) {
	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Get(&ship, getShipByIDQuery, id)
	if err != nil {
		return
	}

	return
}

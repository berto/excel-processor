package db

import _ "github.com/go-sql-driver/mysql"

func RunMigrations() error {
	err := runShipMigration()
	if err != nil {
		return err
	}
	err = runBudgetMigration()
	if err != nil {
		return err
	}
	err = runOpexMigration()
	if err != nil {
		return err
	}
	err = runMonthReportMigration()
	if err != nil {
		return err
	}
	return nil
}

func runShipMigration() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.MustBegin()

	tx.MustExec(createShipTable)

	tx.MustExec("SET FOREIGN_KEY_CHECKS = 0;")
	tx.MustExec(`TRUNCATE TABLE ship;`)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func runBudgetMigration() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.MustBegin()

	tx.MustExec(createBudgetTable)

	tx.MustExec("SET FOREIGN_KEY_CHECKS = 0;")
	tx.MustExec(`TRUNCATE TABLE budget;`)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func runOpexMigration() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.MustBegin()

	tx.MustExec(createOpexTable)

	tx.MustExec("SET FOREIGN_KEY_CHECKS = 0;")
	tx.MustExec(`TRUNCATE TABLE opex;`)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func runMonthReportMigration() error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.MustBegin()

	tx.MustExec(createMonthReportTable)

	tx.MustExec("SET FOREIGN_KEY_CHECKS = 0;")
	tx.MustExec(`TRUNCATE TABLE month_report;`)

	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

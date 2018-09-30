package main

import "log"

const csvPath = "./sample-data/ship.csv"
const csvPathTwo = "./sample-data/ship[1].csv"

func main() {
	err := RunMigrations()
	if err != nil {
		log.Fatal(err)
	}

	err = seedShips(csvPath)
	err = seedShips(csvPathTwo)
	if err != nil {
		log.Fatal(err)
	}
}

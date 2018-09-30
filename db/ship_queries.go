package main

const insertShipQuery = `INSERT INTO ship(
		name, size, year)
		VALUES ( :name, :size, :year)`

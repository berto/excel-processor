package db

const insertShipQuery = `INSERT INTO ship(
		name, size, year)
		VALUES ( :name, :size, :year)`

const getShipByNameQuery = `SELECT * FROM ship WHERE name=?`

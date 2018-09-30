package main

const createShipTable = `CREATE TABLE IF NOT EXISTS ship(
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    name VARCHAR(512),
	    size int,
	    year int,
	    CONSTRAINT ship_pk PRIMARY KEY (id));`

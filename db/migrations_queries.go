package db

const createShipTable = `CREATE TABLE IF NOT EXISTS ship(
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    name VARCHAR(512),
	    size int,
	    year int,
	    CONSTRAINT ship_pk PRIMARY KEY (id));`

const createBudgetTable = `CREATE TABLE IF NOT EXISTS budget(
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    code VARCHAR(512),
	    description VARCHAR(512),
	    CONSTRAINT budget_pk PRIMARY KEY (id));`

const createOpexTable = `CREATE TABLE IF NOT EXISTS opex(
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    ship_id INT UNSIGNED,
	    budget_id INT UNSIGNED,
	    year INT,
	    is_actual BOOL,
	    CONSTRAINT opex_pk PRIMARY KEY (id),
   	 	FOREIGN KEY (ship_id) REFERENCES ship(id) ON DELETE CASCADE,
   	 	FOREIGN KEY (budget_id) REFERENCES budget(id) ON DELETE CASCADE);`

const createMonthReportTable = `CREATE TABLE IF NOT EXISTS month_report(
	    id INT UNSIGNED NOT NULL AUTO_INCREMENT,
	    opex_id INT UNSIGNED,
		month VARCHAR(512), 
		crew_wages FLOAT, 
		crew_expenses FLOAT, 
		shore_based_crew_mgmt FLOAT, 
		victuals FLOAT, 
		insurance_expenses FLOAT, 
		lubricants FLOAT, 
		stores FLOAT, 
		spare_parts FLOAT, 
		repair_and_maintenance FLOAT, 
		other_operating_expenses FLOAT, 
		extraordinary_expenses FLOAT, 
		total FLOAT,
	    CONSTRAINT opex_pk PRIMARY KEY (id),
   	 	FOREIGN KEY (opex_id) REFERENCES opex(id) ON DELETE CASCADE);`

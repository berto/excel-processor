package db

const insertBudgetQuery = `INSERT INTO budget(
		code, description)
		VALUES ( :code, :description)`

const getBudgetByCodeQuery = `SELECT * FROM budget WHERE code=?`

const insertOpexQuery = `INSERT INTO opex(
		ship_id, budget_id, year, is_actual)
		VALUES (:ship_id, :budget_id, :year, :is_actual)`

const getOpexByBudgetIDAndYearQuery = `SELECT * FROM opex WHERE budget_id=? AND year=?`

const getOpexByBudgetIDQuery = `SELECT * FROM opex WHERE budget_id=?`

const getMonthReportsByOpexIDQuery = `SELECT * FROM month_report WHERE opex_id=?`

const insertMonthReportQuery = `INSERT INTO month_report(
			opex_id, 
			month, 
			crew_wages, 
			crew_expenses, 
			shore_based_crew_mgmt, 
			victuals, 
			insurance_expenses, 
			lubricants, 
			stores, 
			spare_parts, 
			repair_and_maintenance, 
			other_operating_expenses, 
			extraordinary_expenses, 
			total)
		VALUES (
			:opex_id, 
			:month, 
			:crew_wages, 
			:crew_expenses, 
			:shore_based_crew_mgmt, 
			:victuals, 
			:insurance_expenses, 
			:lubricants, 
			:stores, 
			:spare_parts, 
			:repair_and_maintenance, 
			:other_operating_expenses, 
			:extraordinary_expenses, 
			:total)`

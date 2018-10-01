package db

type Budget struct {
	ID          int    `db:"id"`
	Code        string `db:"code"`
	Description string `db:"description"`
}

type Opex struct {
	ID       int  `db:"id"`
	ShipID   int  `db:"ship_id"`
	BudgetID int  `db:"budget_id"`
	Year     int  `db:"year"`
	IsActual bool `db:"is_actual"`
}

type MonthReport struct {
	ID                     int     `db:"id" json:"id"`
	OpexID                 int     `db:"opex_id" json:"opex_id,string"`
	Month                  string  `db:"month" json:"month"`
	CrewWages              float32 `db:"crew_wages" json:"crew_wages,string"`
	CrewExpenses           float32 `db:"crew_expenses" json:"crew_expenses,string"`
	ShoreBasedCrewMGMT     float32 `db:"shore_based_crew_mgmt" json:"shore_based_crew_mgmt,string"`
	Victuals               float32 `db:"victuals" json:"victuals,string"`
	InsuranceExpenses      float32 `db:"insurance_expenses" json:"insurance_expenses,string"`
	Lubricants             float32 `db:"lubricants" json:"lubricants,string"`
	Stores                 float32 `db:"stores" json:"stores,string"`
	SpareParts             float32 `db:"spare_parts" json:"spare_parts,string"`
	RepairAndMaintenance   float32 `db:"repair_and_maintenance" json:"repair_and_maintenance,string"`
	OtherOperatingExpenses float32 `db:"other_operating_expenses" json:"other_operating_expenses,string"`
	ExtraordinaryExpenses  float32 `db:"extraordinary_expenses" json:"extraordinary_expenses,string"`
	Total                  float32 `db:"total" json:"total,string"`
}

func AddBudget(budget Budget) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	tx := db.MustBegin()

	_, err = tx.NamedExec(insertBudgetQuery, budget)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	err = db.Get(&budget, getBudgetByCodeQuery, budget.Code)
	if err != nil {
		return 0, err
	}

	return budget.ID, nil
}

func AddOpex(opex Opex) (int, error) {
	db, err := getDB()
	if err != nil {
		return 0, err
	}
	defer db.Close()

	tx := db.MustBegin()

	_, err = tx.NamedExec(insertOpexQuery, opex)
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	err = db.Get(&opex, getOpexByBudgetIDQuery, opex.BudgetID)
	if err != nil {
		return 0, err
	}

	return opex.ID, nil
}

func AddMonthReport(monthReport MonthReport) error {
	db, err := getDB()
	if err != nil {
		return err
	}
	defer db.Close()

	tx := db.MustBegin()

	_, err = tx.NamedExec(insertMonthReportQuery, monthReport)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

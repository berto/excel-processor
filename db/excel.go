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

type BudgetDataResponse struct {
	Error string `json:"error"`
	DataMessage
}

type DataMessage struct {
	Code     string     `json:"code"`
	OpexData []OpexData `json:"opexData"`
}

type OpexData struct {
	ShipName     string        `json:"shipName"`
	Year         int           `json:"year"`
	IsActual     bool          `json:"isActual"`
	MonthReports []MonthReport `json:"monthReports"`
}

func GetBudgetDataFromCode(code string) (br BudgetDataResponse) {
	db, err := getDB()
	if err != nil {
		br.Error = err.Error()
		return
	}
	defer db.Close()

	budget := Budget{}
	err = db.Get(&budget, getBudgetByCodeQuery, code)
	if err != nil {
		br.Error = err.Error()
		return
	}
	dm := DataMessage{}
	dm.Code = code

	opexs, err := GetOpexsByBudgetID(budget.ID)
	if err != nil {
		br.Error = err.Error()
		return
	}

	dm.OpexData = make([]OpexData, len(opexs))

	for i, opex := range opexs {
		ship, err := FindShipByID(opex.ShipID)
		if err != nil {
			br.Error = err.Error()
			return
		}
		monthReports, err := GetMonthReportsByOpexID(opex.ID)
		if err != nil {
			br.Error = err.Error()
			return
		}
		od := OpexData{
			ShipName:     ship.Name,
			Year:         opex.Year,
			IsActual:     opex.IsActual,
			MonthReports: monthReports,
		}
		dm.OpexData[i] = od
	}

	br.DataMessage = dm
	return
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

	err = db.Get(&opex, getOpexByBudgetIDAndYearQuery, opex.BudgetID, opex.Year)
	if err != nil {
		println(err.Error())
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

func GetOpexsByBudgetID(budgetID int) (opexs []Opex, err error) {
	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Select(&opexs, getOpexByBudgetIDQuery, budgetID)
	if err != nil {
		return
	}
	return
}

func GetMonthReportsByOpexID(opexID int) (monthReports []MonthReport, err error) {
	db, err := getDB()
	if err != nil {
		return
	}
	defer db.Close()

	err = db.Select(&monthReports, getMonthReportsByOpexIDQuery, opexID)
	if err != nil {
		return
	}
	return
}

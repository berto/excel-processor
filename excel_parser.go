package main

import (
	"encoding/json"
	"strconv"
	"strings"
	"unicode"

	"github.com/berto/excel-processor/db"
	"github.com/tealeg/xlsx"
)

var categories = map[string]string{
	"CREW WAGES Total":               "crew_wages",
	"CREW EXPENSES Total":            "crew_expenses",
	"SHORE-BASED CREW MGMT Total":    "shore_based_crew_mgmt",
	"VICTUALS Total":                 "victuals",
	"INSURANCE EXPENSES Total":       "insurance_expenses",
	"LUBRICANTS Total":               "lubricants",
	"SHORES Total":                   "stores",
	"SPARE PARTS Total":              "spare_parts",
	"REPAIR & MAINTENANCE Total":     "repair_and_maintenance",
	"OTHER OPERATING EXPENSES Total": "other_operating_expenses",
	"EXTRAORDINARY EXPENSES Total":   "extraordinary_expenses",
	"Grand Total":                    "total",
}

var monthNames = []string{"January", "February", "March", "April", "May", "June", "July", "August", "October", "September", "November", "December"}

func parseExcelData(xlFile *xlsx.File, name string) error {
	shipName, code := parseName(name)
	ship, err := db.FindShipByName(shipName)
	if err != nil {
		return err
	}
	budget := db.Budget{Code: code}
	budgetID, err := db.AddBudget(budget)
	if err != nil {
		return err
	}
	for _, sheet := range xlFile.Sheets {
		skip, err := parseSheetData(sheet, ship.ID, budgetID)
		if err != nil {
			return err
		}
		if skip {
			break
		}
	}
	return nil
}

func parseSheetData(sheet *xlsx.Sheet, shipID int, budgetID int) (bool, error) {
	year, isActual := parseSheetName(sheet.Name)
	if year == 0 {
		return true, nil
	}
	opex := db.Opex{
		ShipID:   shipID,
		BudgetID: budgetID,
		Year:     year,
		IsActual: isActual,
	}
	opexID, err := db.AddOpex(opex)
	if err != nil {
		return false, err
	}

	months := make([]map[string]string, 12)
	for _, row := range sheet.Rows {
		parseRowData(row, months)
	}
	for i, month := range months {
		err = createMonthReport(i, opexID, month)
		if err != nil {
			return false, err
		}
	}
	return false, nil
}

func parseRowData(row *xlsx.Row, months []map[string]string) {
	totalRows := 14
	var category string
	for i, cell := range row.Cells {
		if i == 1 {
			if categories[cell.String()] == "" {
				break
			}
			category = categories[cell.String()]
		}
		if i > 1 && i < totalRows {
			if months[i-2] == nil {
				months[i-2] = make(map[string]string)
			}
			months[i-2][category] = cell.String()
		}
	}
}

func createMonthReport(i int, opexID int, month map[string]string) error {
	if month["total"] == "" {
		return nil
	}
	report := db.MonthReport{}
	month["opex_id"] = strconv.Itoa(opexID)
	month["month"] = monthNames[i]
	monthJSON, err := json.Marshal(month)
	if err != nil {
		return err
	}
	err = json.Unmarshal([]byte(monthJSON), &report)
	if err != nil {
		return err
	}
	err = db.AddMonthReport(report)
	if err != nil {
		return err
	}
	return nil
}

func parseName(name string) (shipName, code string) {
	split := strings.SplitAfter(name, "-")
	seperateName := strings.SplitAfter(split[1], "(")
	code = split[0][:3]
	shipName = toSnake(seperateName[0][1 : len(seperateName[0])-2])
	shipName = strings.TrimRight(shipName, "\n")
	code = strings.TrimRight(code, "\n")
	return
}

func parseSheetName(name string) (year int, isActual bool) {
	split := strings.SplitAfter(name, " ")
	yearString := strings.Trim(split[0], " ")
	year, err := strconv.Atoi(yearString)
	if err != nil || len(split) < 2 {
		return
	}
	return year, split[1] == "Actual"
}

func toSnake(in string) string {
	runes := []rune(in)
	length := len(runes)

	var out []rune
	for i := 0; i < length; i++ {
		if i > 0 && unicode.IsUpper(runes[i]) && ((i+1 < length && unicode.IsLower(runes[i+1])) || unicode.IsLower(runes[i-1])) {
			out = append(out, '_')
		}
		if runes[i] != ' ' {
			out = append(out, unicode.ToLower(runes[i]))
		} else {
			out = append(out, '_')
		}
	}

	return string(out)
}

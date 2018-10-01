package main

import "testing"

func TestExcelParse(t *testing.T) {
	fileName := "084 - SEASPAN HAMBURG (2015A-2018F).xlsx"
	fileNameTwo := "304 - CSCL PANAMA (2015A-2018F).xlsx"
	name, code := parseName(fileName)
	nameTwo, codeTwo := parseName(fileNameTwo)

	if name != "seaspan_hamburg" || code != "084" {
		t.Errorf("Failed to parse file name %s, %s", name, code)
	}
	if nameTwo != "cscl_panama" || codeTwo != "304" {
		t.Errorf("Failed to parse file name two %s, %s", nameTwo, codeTwo)
	}

	sheetName := "2015 Actual"
	sheetNameTwo := "2018 Forecast"
	sheetNameThree := "SSR"
	year, isActual := parseSheetName(sheetName)
	yearTwo, isActualTwo := parseSheetName(sheetNameTwo)
	yearThree, isActualThree := parseSheetName(sheetNameThree)

	if year != 2015 || isActual != true {
		t.Errorf("Failed to parse sheet name %v, %v", year, isActual)
	}
	if yearTwo != 2018 || isActualTwo != false {
		t.Errorf("Failed to parse sheet name two %v, %v", yearTwo, isActualTwo)
	}
	if yearThree != 0 || isActualThree != false {
		t.Errorf("Failed to parse sheet name three %v, %v", yearThree, isActualThree)
	}
}

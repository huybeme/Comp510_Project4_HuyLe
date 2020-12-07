package main

import (
	"os"
	"testing"
)

func TestCreateTables(t *testing.T) {

	emptyTables := createTables(GameDataBase)
	if len(emptyTables) == 0 {
		t.Error("expected an array of headers and got 0", emptyTables)
	}

	insufficientTables := createTables(GameDataBase)
	if len(insufficientTables) != sheet.MaxCol+1 {
		t.Error("did not extract all columns from excel sheet", insufficientTables)
	}

	if notexist, err := os.Stat(".gamedatabase.db"); os.IsExist(err) {
		t.Error("game database file does not exist", notexist)
	}

}

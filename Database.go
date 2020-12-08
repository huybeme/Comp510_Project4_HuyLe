package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"github.com/tealeg/xlsx/v3"
	"strconv"
)

const (
	path      = "./games-features.xlsx"
	mainSheet = "games-features"
)

var (
	data, _      = xlsx.OpenFile(path)
	sheet, _     = data.Sheet[mainSheet]
	GameDataBase *sql.DB
)

func OpenDatabase(dbfile string) *sql.DB {
	database, err := sql.Open("sqlite3", dbfile)
	CheckErr("Open database error", err)
	fmt.Println("GameData database opened")
	return database
}

func fillDatabase(db *sql.DB) {
	headers := createTables(db)
	var keyStatement string
	var updateStatement string
	var cellElement *xlsx.Cell

	id := 1
	for i := -1; i < sheet.MaxCol; i++ {
		keyStatement = "INSERT INTO GameData (" + headers[i+1] + ") VALUES (?)"
		prepKey, err := db.Prepare(keyStatement)
		CheckErr("fillDatabase: PrepStatement error ", err)
		updateStatement = "UPDATE GameData SET " + headers[i+1] + " = ? WHERE id = ?"
		prepUpdate, err := db.Prepare(updateStatement)

		for j := 1; j < sheet.MaxRow; j++ {
			if i == -1 {
				_, err = prepKey.Exec(id)
				CheckErr("fillDatabase: PrepStatement error at if statement ", err)
				id++
			} else {
				cellElement, err = sheet.Cell(j, i)
				CheckErr("fillDatabase: Retrieve cell value error ", err)
				_, err = prepUpdate.Exec(cellElement.Value, id)
				CheckErr("fillDatabase: PrepStatement error at else statement ", err)
				id++
			}
			if j%500 == 0 {
				fmt.Println("\t data appended at row", strconv.Itoa(j), "/", strconv.Itoa(sheet.MaxRow), "complete")
			}
		} // end of inner for loop
		fmt.Println("data appended to column '" + headers[i+1] + "' complete (" + strconv.Itoa(i+1) + "/" + strconv.Itoa(sheet.MaxCol) + ")")
		id = 1
	} // end outer for loop
	fmt.Println("fill database with values complete")
}

func createTables(db *sql.DB) []string {
	statement := "CREATE TABLE IF NOT EXISTS GameData(" +
		"id INTEGER PRIMARY KEY,\n"
	headerType := ""
	var cellElement *xlsx.Cell
	var cellType *xlsx.Cell
	var headers []string
	headers = append(headers, "id")

	for i := 0; i < sheet.MaxCol; i++ {

		cellType, _ = sheet.Cell(1, i)
		cellElement, _ = sheet.Cell(0, i)
		if i == sheet.MaxCol-1 {
			if cellType.Type() == xlsx.CellTypeNumeric {
				headerType = "INTEGER"
			} else if cellType.Type() == xlsx.CellTypeBool {
				headerType = "BOOLEAN"
			} else {
				headerType = "TEXT"
			}
			statement = statement + cellElement.Value + " " + headerType + ");"
		} else {
			if cellType.Type() == xlsx.CellTypeNumeric {
				headerType = "INTEGER"
			} else if cellType.Type() == xlsx.CellTypeBool {
				headerType = "BOOLEAN"
			} else {
				headerType = "TEXT"
			}
			statement = statement + cellElement.Value + " " + headerType + ",\n"
		}
		headers = append(headers, cellElement.Value)
	}

	//if isGameDataEmpty() {
	//_, err := db.Exec(statement)
	//if err != nil {
	//	log.Fatal("Database execution error ", err)
	//}
	//}
	fmt.Println("Database Tables Created")
	return headers
}

//func isGameDataEmpty() bool {
//	result, err := GameDataBase.Query("SELECT id FROM GameData")
//	CheckErr("Query error at isTableEmpty ", err)
//	defer result.Close()
//
//	count := 0
//	for result.Next() {
//		count++
//	}
//
//	if count > 0 {
//		return false
//	}
//	return true
//}

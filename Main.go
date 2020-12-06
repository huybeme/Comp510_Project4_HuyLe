package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"reflect"
	"strings"
)

func CheckErr(prompt string, err error) {
	if err != nil {
		log.Fatal(prompt, err)
	}
}

type Entry struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func main() {

	GameDataBase = OpenDatabase(".GameData.db")
	defer GameDataBase.Close()
	//fillDatabase(GameDataBase)

	router := mux.NewRouter()

	//router.HandleFunc("/gameapi", getEntries).Methods("GET")
	router.HandleFunc("/gameapi", getEntries).Methods("GET")
	//router.HandleFunc("/gameapi", createEntry).Methods("POST")
	router.HandleFunc("/gameapi/{id}", getEntry).Methods("GET")
	//router.HandleFunc("/gameapi/{id}", udateEntry).Methods("PUT")
	//router.HandleFunc("/gameapi/{api}", deleteEntry).Methods("DELETE")

	router.HandleFunc("/gameapi/search/", searchEntries).Methods("GET")

	port := ":8080"
	err := http.ListenAndServe(port, router)
	CheckErr("unable to connect to port '"+port+"' ", err)

}

func searchEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	result, err := GameDataBase.Query("SELECT id, QueryName FROM GameData")
	CheckErr("Query error at searchEntries ", err)
	defer result.Close()

	var entries []Entry
	for result.Next() {
		var entry Entry
		err = result.Scan(&entry.ID, &entry.Title)
		CheckErr("Scan error at getEntries ", err)

		fmt.Println(reflect.TypeOf(entry.Title))

		if strings.Contains(entry.Title, "half") {
			entries = append(entries, entry)
		}
	}
	fmt.Println(entries)

}

func getEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var entries []Entry
	result, err := GameDataBase.Query("SELECT id, QueryName FROM GameData")
	CheckErr("Query error at getEntries ", err)
	defer result.Close()

	for result.Next() {
		var entry Entry
		err = result.Scan(&entry.ID, &entry.Title)
		CheckErr("Scan error at getEntries ", err)
		entries = append(entries, entry)
	}
	json.NewEncoder(w).Encode(entries)
	fmt.Println("gamedata API entries accessed")
}

func getEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := mux.Vars(r)
	result, err := GameDataBase.Query("SELECT id, QueryName FROM GameData WHERE id = ?", key["id"])
	CheckErr("Query error at getEntry ", err)
	defer result.Close()

	var entry Entry
	for result.Next() {
		err = result.Scan(&entry.ID, &entry.Title)
		CheckErr("Scan error at getEntries ", err)
	}
	json.NewEncoder(w).Encode(entry)
	fmt.Println("gamedata API entry accessed")
}

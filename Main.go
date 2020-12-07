package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

func CheckErr(prompt string, err error) {
	if err != nil {
		log.Fatal(prompt, err)
	}
}

type Entry struct {
	ID                  int     `json:"id"`
	Title               string  `json:"title"`
	ReleaseDate         string  `json:"release_date"`
	RecommendationCount int     `json:"recommendation_count"`
	Price               float64 `json:"price"`
	Description         string  `json:"desc"`
	PCminReq            string  `json:"pc_min_req"`
	LinuxMinReq         string  `json:"linux_min_req"`
	MacMinReq           string  `json:"mac+min_req"`
}

func main() {

	GameDataBase = OpenDatabase(".GameData.db")
	defer GameDataBase.Close()
	//fillDatabase(GameDataBase)

	router := mux.NewRouter()

	router.HandleFunc("/gameapi/all-entries/", getEntries).Methods("GET")
	router.HandleFunc("/gameapi/single-entry/{id}", getEntry).Methods("GET")

	// functions not created, may be out of the scope of this project since our database source is an excel file
	//router.HandleFunc("/gameapi", createEntry).Methods("POST")
	//router.HandleFunc("/gameapi/{id}", udateEntry).Methods("PUT")
	//router.HandleFunc("/gameapi/{api}", deleteEntry).Methods("DELETE")

	router.HandleFunc("/gameapi", getHomepage)
	router.HandleFunc("/gameapi/query/{search}", searchEntries).Methods("GET")

	port := ":8080"
	err := http.ListenAndServe(port, router)
	CheckErr("unable to connect to port '"+port+"' ", err)

}

type Homepage struct {
	Title string
	News  string
}

func getHomepage(w http.ResponseWriter, r *http.Request) {
	header := Homepage{"Welcome to the GameAPI Homepage", "Perform a search below."}
	temp, err := template.ParseFiles("./index.html")
	CheckErr("Template parse files error at getHomepage ", err)
	err = temp.Execute(w, header)
	CheckErr("Template execute error at getHomepage", err)

	fmt.Println("Homepage Reached")
}

func searchEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	query := mux.Vars(r)
	result, err := GameDataBase.Query("SELECT " +
		"id, QueryName, ReleaseDate, RecommendationCount, PriceFinal, AboutText, PCMinReqsText, LinuxMinReqsText, " +
		"MacMinReqsText FROM GameData WHERE QueryName LIKE '%" + query["search"] + "%'")
	CheckErr("Query error at searchEntries ", err)
	defer result.Close()

	var entries []Entry
	for result.Next() {
		var entry Entry
		err = result.Scan(&entry.ID, &entry.Title, &entry.ReleaseDate, &entry.RecommendationCount, &entry.Price,
			&entry.Description, &entry.PCminReq, &entry.LinuxMinReq, &entry.MacMinReq)
		CheckErr("Scan error at getEntries ", err)
		entries = append(entries, entry)
	}
	json.NewEncoder(w).Encode(entries)
	fmt.Println("Search entry complete")
}

func getEntries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var entries []Entry
	result, err := GameDataBase.Query("SELECT " +
		"id, QueryName, ReleaseDate, RecommendationCount, PriceFinal, AboutText, PCMinReqsText, LinuxMinReqsText, " +
		"MacMinReqsText FROM GameData")
	CheckErr("Query error at getEntries ", err)
	defer result.Close()

	for result.Next() {
		var entry Entry
		err = result.Scan(&entry.ID, &entry.Title, &entry.ReleaseDate, &entry.RecommendationCount, &entry.Price,
			&entry.Description, &entry.PCminReq, &entry.LinuxMinReq, &entry.MacMinReq)
		CheckErr("Scan error at getEntries ", err)
		entries = append(entries, entry)
	}
	json.NewEncoder(w).Encode(entries)
	fmt.Println("gamedata API entries accessed")
}

func getEntry(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	key := mux.Vars(r)
	result, err := GameDataBase.Query("SELECT "+
		"id, QueryName, ReleaseDate, RecommendationCount, PriceFinal, AboutText, PCMinReqsText, LinuxMinReqsText, "+
		"MacMinReqsText FROM GameData WHERE id = ?", key["id"])
	CheckErr("Query error at getEntry ", err)
	defer result.Close()

	var entry Entry
	for result.Next() {
		err = result.Scan(&entry.ID, &entry.Title, &entry.ReleaseDate, &entry.RecommendationCount, &entry.Price,
			&entry.Description, &entry.PCminReq, &entry.LinuxMinReq, &entry.MacMinReq)
		CheckErr("Scan error at getEntries ", err)
	}
	json.NewEncoder(w).Encode(entry)
	fmt.Println("gamedata API entry accessed")
}

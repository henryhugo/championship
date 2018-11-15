package main

import (
	"fmt"
	"net/http"
	"os"
)

type LeaguesStorage interface {
	Init()
	Add(s League) error
}

var Global_db LeaguesStorage

type League struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	ID      string `json:"id"`
}

type LeagueDB struct {
	leagues map[string]League
}

func (db *LeagueDB) Init() {
	db.leagues = make(map[string]League)
}
func (db *LeagueDB) Add(l League) error {
	db.leagues[l.ID] = l
	return nil
}

func league(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func champ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api for championship league")

}

var db LeagueDB

func main() {

	db = LeagueDB{}

	//Global_db.Init()
	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/league", league)
	http.ListenAndServe(":"+port, nil)
}

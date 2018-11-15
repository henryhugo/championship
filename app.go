package main

import (
	"championship/leaguedb"
	"fmt"
	"net/http"
	"os"
)

func league(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func champ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api for championship league")

}

func main() {

	//Global_db.Init()
	leaguedb.Global_db = &leaguedb.LeaguesDB{}
	leaguedb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/league", league)
	http.ListenAndServe(":"+port, nil)
}

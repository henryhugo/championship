package main

import (
	"championship/leaguedb"
	"fmt"
	"net/http"
	"os"
)

func champ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "api for championship league")

}

func main() {
	//in memory strorage
	leaguedb.Global_db = &leaguedb.LeaguesDB{}

	//mongodb storage
	leaguedb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/league", leaguedb.LeaguesStorage.leagueHandler)
	http.ListenAndServe(":"+port, nil)
}

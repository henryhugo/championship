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
	leaguedb.InitWh()
	//in memory strorage
	leaguedb.Global_db = &leaguedb.LeaguesMongoDB{
		DatabaseURL:           "mongodb://hugoh:6926a5b8@ds057548.mlab.com:57548/championship",
		DatabaseName:          "championship",
		LeaguesCollectionName: "league",
	}


	//mongodb storage
	leaguedb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/leagues/", leaguedb.LeagueHandler)  //POST
	http.HandleFunc("/champ/webhook/", leaguedb.WebhookHandler) //POST et GET
	//http.HandleFunc("/champ/(nomdepays)/(date)||(nomequipe)||(day)", leaguedb.MatchHandler)
	http.ListenAndServe(":"+port, nil)
}

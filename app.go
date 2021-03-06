package main

import (
	"championship/leaguedb"
	"championship/matchdb"
	"fmt"
	"net/http"
	"os"
)

func champ(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API for championship league")
}

func main() {
	leaguedb.InitWh()
	matchdb.InitWh()

	leaguedb.Global_db = &leaguedb.LeaguesMongoDB{
		DatabaseURL:           "mongodb://hugoh:@ds057548.mlab.com:57548/championship",
		DatabaseName:          "championship",
		LeaguesCollectionName: "league",
	}

	leaguedb.Global_db.Init()

	matchdb.Global_db = &matchdb.MatchesMongoDB{
		DatabaseURL:           "mongodb://hugoh:@ds057548.mlab.com:57548/championship",
		DatabaseName:          "championship",
		MatchesCollectionName: "matchs",
	}

	matchdb.Global_db.Init()

	port := os.Getenv("PORT")

	http.HandleFunc("/champ", champ)
	http.HandleFunc("/champ/leagues/", leaguedb.LeagueHandler)              //POST and GET
	http.HandleFunc("/champ/webhookLeague/", leaguedb.WebhookLeagueHandler) //POST and GET
	http.HandleFunc("/champ/matchs/", matchdb.MatchHandler)
	http.HandleFunc("/champ/webhookMatch/", matchdb.WebhookMatchHandler)

	http.ListenAndServe(":"+port, nil)
}

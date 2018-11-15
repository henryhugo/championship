package main

import (
	. "championship/dao"
	. "championship/models"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

//var config = Config{}
var dao = LeaguesDAO{}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllLeaguesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func FindLeagueEndpoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func CreateLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var league League
	if err := json.NewDecoder(r.Body).Decode(&league); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	league.CountryID = bson.NewObjectId()
	//league.CountryName = "England"
	league.LeagueID = bson.NewObjectId()
	//league.LeagueName = "Premier League"
	if err := dao.Insert(league); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, league)
}

func UpdateLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

func DeleteLeagueEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "not implemented yet !")
}

/*func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}*/

func main() {
	/*addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}*/
	port := os.Getenv("PORT")
	r := mux.NewRouter()
	r.HandleFunc("/league", AllLeaguesEndPoint).Methods("GET")
	r.HandleFunc("/league", CreateLeagueEndPoint).Methods("POST")
	r.HandleFunc("/league", UpdateLeagueEndPoint).Methods("PUT")
	r.HandleFunc("/league", DeleteLeagueEndPoint).Methods("DELETE")
	r.HandleFunc("/league/{id}", FindLeagueEndpoint).Methods("GET")
	http.ListenAndServe(":"+port, nil)
	//log.Fatal(http.ListenAndServe(addr, nil))

	/*if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}*/
}

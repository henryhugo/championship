package leaguedb

import (
	"net/http"
)

type LeaguesStorage interface {
	Init()
	Add(l League) error
	leagueHandler(w http.ResponseWriter, r *http.Request)
}

type League struct {
	Name     string `json:"name"`
	Country  int    `json:"country"`
	LeagueID string `json:"leagueid"`
}

type LeaguesDB struct {
	leagues map[string]League
}

func (db *LeaguesDB) Init() {
	db.leagues = make(map[string]League)
}

func (db *LeaguesDB) Add(l League) error {
	db.leagues[l.LeagueID] = l
	return nil
}

/*func leagueHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		var l League
		err := json.NewDecoder(r.Body).Decode(&l)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		Global_db.Add(l)
		fmt.Fprint(w, "ok") // 200 by default
		return

	case "GET":
		fmt.Fprintln(w, "not implemented yet")

	default:

		fmt.Fprintln(w, "not implemented yet !")
	}
}*/

package leaguedb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func LeagueHandler(w http.ResponseWriter, r *http.Request) {
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
}

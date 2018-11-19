package leaguedb

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type webhook struct {
	WebhookURL string //"webhookURL": "http://remoteUrl:8080/randomWebhookPath",
}

var whDB map[string]webhook
var idCountwh int

func LeagueHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			var l League
			err := json.NewDecoder(r.Body).Decode(&l)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			_, ok := Global_db.Get(l.LeagueID)
			if ok {
				// TODO find a better Error Code (HTTP Status)
				http.Error(w, "League already exists.", http.StatusBadRequest)
				return
			}
			Global_db.Add(l)
			text := "{\"text\": \"New league added to the database :" + l.Name + " !\"}"
			payload := strings.NewReader(text)
			for _, wh := range whDB {
				client := &http.Client{Timeout: (time.Second * 30)}
				req, err := http.NewRequest("POST", wh.WebhookURL, payload)
				req.Header.Set("Content-Type", "application/json")
				resp, err := client.Do(req)
				if err != nil {
					fmt.Print(err.Error())
				}
				fmt.Println(resp.Status)
				fmt.Fprint(w, "ok") 
				return
			}
		}

	case "GET":
		http.Header.Add(w.Header(), "content-type", "application/json")
		parts := strings.Split(r.URL.Path, "/")
		switch {
		case pathLeague.MatchString(r.URL.Path):
			{

				fmt.Fprint(w, Global_db.DisplayLeague())
			}
		case pathidfield.MatchString(r.URL.Path):
			{
				var l League
				id := parts[3]
				infoWanted := parts[4]
				l, ok := Global_db.Get(id)
				if !ok {
					// TODO find a better Error Code (HTTP Status)
					http.Error(w, "League don't exists.", http.StatusBadRequest)
					return
				}
				switch infoWanted {
				case "country":
					{
						fmt.Fprint(w, l.Country)
					}
				case "leagueID":
					{
						fmt.Fprint(w, l.LeagueID)
					}
				case "name":
					{
						fmt.Fprint(w, l.Name)
					}
				case "teams":
					{
						fmt.Fprint(w, l.Teams)
					}
				default:
					fmt.Fprint(w, "Not found")

				}
			}
		case pathid.MatchString(r.URL.Path):
			{
				id := parts[3]
				l, ok := Global_db.Get(id)
				if !ok {
					// TODO find a better Error Code (HTTP Status)
					http.Error(w, "League don't exists.", http.StatusBadRequest)
					return
				}
				json.NewEncoder(w).Encode(l)
			}

		}

	default:

		fmt.Fprintln(w, "not implemented yet !")

	}
}

func WebhookLeagueHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "POST":
		{
			var wh webhook
			//TODO check correct wh format
			err := json.NewDecoder(r.Body).Decode(&wh)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
			}
			Idstr := "id"
			strValue := fmt.Sprintf("%d", idCountwh)
			newId := Idstr + strValue
			idCountwh += 1
			whDB[newId] = wh
			json.NewEncoder(w).Encode(newId)

		}
	case "GET":
		{
			if pathwhID.MatchString(r.URL.Path) {
				idWant := parts[3]
				for id, file := range whDB {
					if id == idWant {
						json.NewEncoder(w).Encode(file)
					}
				}

			}

		}
	}
}

func InitWh() {
	idCountwh = 0
	whDB = map[string]webhook{}

}

var pathwhID, _ = regexp.Compile("/champ/webhook/id[0-9]+$")
var pathLeague, _ = regexp.Compile("/champ/leagues[/]{1}$")
var pathidfield, _ = regexp.Compile("/champ/leagues/id[0-9]+/(country$|name$|leagueID$|teams$)")
var pathid, _ = regexp.Compile("/champ/leagues/id[0-9]+$")

package matchdb

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

func MatchHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		{
			var m MatchesL
			err := json.NewDecoder(r.Body).Decode(&m)
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			_, ok := Global_db.Get(m.LeagueID)
			if ok {
				// TODO find a better Error Code (HTTP Status)
				http.Error(w, "Matches of this League already exists.", http.StatusBadRequest)
				return
			}
			Global_db.Add(m)
			text := "{\"text\": \"New matchs added to the database :" + m.Name + " !\"}"
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
				fmt.Fprint(w, "ok") // 200 by default
				return
			}
		}

	case "GET":
		http.Header.Add(w.Header(), "content-type", "application/json")
		//parts := strings.Split(r.URL.Path, "/")
		switch {
		case pathMatches.MatchString(r.URL.Path):
			{

				//fmt.Fprint(w, Global_db.DisplayMatches())
			}
			/*case pathTeamid.MatchString(r.URL.Path):
			{
				teamName := parts[4]
				fmt.Fprint(w, Global_db.FindTeam(teamName))
			}*/

		}

	default:

		fmt.Fprintln(w, "not implemented yet !")

	}
}

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")
	parts := strings.Split(r.URL.Path, "/")
	switch r.Method {
	case "POST":
		{
			//fmt.Fprintln(w, "wh")
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
var pathMatches, _ = regexp.Compile("/champ/matches/[a-z]{3}$")

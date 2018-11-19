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
		parts := strings.Split(r.URL.Path, "/")
		switch {
		case pathMatchs.MatchString(r.URL.Path):
			{
				fmt.Fprint(w, Global_db.DisplayMatches())
			}
		case pathMatchsID.MatchString(r.URL.Path):
			{
				id := parts[3]
				m, ok := Global_db.Get(id)
				if !ok {
					// TODO find a better Error Code (HTTP Status)
					http.Error(w, "Matchs don't exists.", http.StatusBadRequest)
					return
				}
				json.NewEncoder(w).Encode(m)
			}
		case pathMatchsFields.MatchString(r.URL.Path):
			{
				var m MatchesL
				id := parts[3]
				infoWanted := parts[4]
				m, ok := Global_db.Get(id)
				if !ok {
					// TODO find a better Error Code (HTTP Status)
					http.Error(w, "Matchs don't exists.", http.StatusBadRequest)
					return
				}
				switch infoWanted {
				case "name":
					fmt.Fprint(w, m.Name)
				case "leagueID":
					fmt.Fprint(w, m.LeagueID)
				case "rounds":
					fmt.Fprint(w, m.Rounds)
				default:
					fmt.Fprint(w, "Not found")

				}

			}
		case pathMatchday.MatchString(r.URL.Path): //matchdayX
			{
				var m MatchesL
				id := parts[3]
				infoWanted := parts[4]
				infoWanted = strings.Replace(infoWanted, "m", "M", 1)
				tab := strings.SplitAfterN(infoWanted, "y", 2)

				info := tab[0] + " " + tab[1]
				fmt.Fprint(w, info)
				m, ok := Global_db.Get(id)
				if !ok {
					// TODO find a better Error Code (HTTP Status)
					http.Error(w, "Matchs don't exists.", http.StatusBadRequest)
					return
				}

				for _, r := range m.Rounds {
					if r.Name == info {
						json.NewEncoder(w).Encode(r.Matches)
					}
				}

			}
		case pathMatchFields.MatchString(r.URL.Path): //
			{
				var m MatchesL
				id := parts[3]
				infoWanted := parts[4]
				infoWanted = strings.Replace(infoWanted, "m", "M", 1)
				tab := strings.SplitAfterN(infoWanted, "y", 2)
				iw := tab[0] + " " + tab[1]
				info := parts[5]
				m, ok := Global_db.Get(id)
				if !ok {
					// TODO find a better Error Code (HTTP Status)
					http.Error(w, "Matchs don't exists.", http.StatusBadRequest)
					return
				}
				fmt.Fprint(w, info+" "+iw)
				switch info {
				case "date":
					{
						for _, r := range m.Rounds {
							if r.Name == iw {
								for _, m := range r.Matches {
									json.NewEncoder(w).Encode(m.Date)
								}
							}
						}
					}
				case "team1":
					{

					}
				case "team2":
					{

					}
				case "score1":
					{

					}
				case "score2":
					{

					}
				}
			}

		}

	default:

		fmt.Fprintln(w, "not implemented yet !")

	}
}

func WebhookMatchHandler(w http.ResponseWriter, r *http.Request) {
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
var pathMatchs, _ = regexp.Compile("/champ/matchs/$")
var pathMatchsID, _ = regexp.Compile("/champ/matchs/id[0-9]+$")
var pathMatchday, _ = regexp.Compile("/champ/matchs/id[0-9]+/matchday[0-9]+$")
var pathMatchsFields, _ = regexp.Compile("/champ/matchs/id[0-9]+/(name$|leagueID$|rounds$)")
var pathMatchFields, _ = regexp.Compile("/champ/matchs/id[0-9]+/matchday[0-9]+/(date$|team1$|team2$|score1|score2)")

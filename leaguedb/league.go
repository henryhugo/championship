package leaguedb

type LeaguesStorage interface {
	Init()
	Add(l League) error
	Get(key string) (League, bool)
	/*DisplayLeagueName() string
	FindTeam(team string) string*/
}

type League struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
	LeagueID string `json:"leagueid"`
	Teams    []Team `json:"teams"`
}

type Team struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Code string `json:"code"`
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

func (db *LeaguesDB) Get(keyID string) (League, bool) {
	l, ok := db.leagues[keyID]
	return l, ok
}

func (db *LeaguesDB) DisplayLeagueName() string {
	str := ""
	for _, l := range db.leagues {
		str = str + l.Name
	}
	return str

}

func (db *LeaguesDB) findTeam(team string) string {
	str := "Your team is not in the database"
	for _, l := range db.leagues {
		for _, t := range l.Teams {
			if t.Name == team {
				str := "Your team play in league " + l.Name + "their code is" + t.Code
				return str
			}
		}
	}
	return str
}

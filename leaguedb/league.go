package leaguedb

type LeaguesStorage interface {
	Init()
	Add(s League) error
}

type League struct {
	Name    string `json:"name"`
	Country int    `json:"country"`
	ID      string `json:"id"`
}

type LeaguesDB struct {
	leagues map[string]League
}

func (db *LeaguesDB) Init() {
	db.leagues = make(map[string]League)
}

func (db *LeaguesDB) Add(l League) error {
	db.leagues[l.ID] = l
	return nil
}

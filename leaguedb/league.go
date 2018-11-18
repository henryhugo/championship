package leaguedb

type LeaguesStorage interface {
	Init()
	Add(l League) error
	Get(key string) (League, bool)
}

type League struct {
	Name     string `json:"name"`
	Country  string `json:"country"`
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

func (db *LeaguesDB) Get(keyID string) (League, bool) {
	l, ok := db.leagues[keyID]
	return l, ok
}

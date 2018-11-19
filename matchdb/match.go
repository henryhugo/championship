package matchdb

type MatchesStorage interface {
	Init()
	Add(m MatchesL) error
	Get(key string) (MatchesL, bool)
	DisplayMatches() MatchesL
}

type Team struct {
	Key  string `json:"key"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type Match struct {
	Date   string `json:"date"`
	Team1  Team   `json:"team1"`
	Team2  Team   `json:"team2"`
	Score1 int64  `json:"score1"`
	Score2 int64  `json:"score2"`
}

type Round struct {
	Name    string  `json:"name"`
	Matches []Match `json:"matches"`
}

type MatchesL struct {
	Name     string  `json:"name"`
	LeagueID string  `json:"leagueid"`
	Rounds   []Round `json:"rounds"`
}

type MatchesLDB struct {
	matchesL map[string]MatchesL
}

func (db *MatchesLDB) Init() {
	db.matchesL = make(map[string]MatchesL)
}

func (db *MatchesLDB) Add(m MatchesL) error {
	db.matchesL[m.LeagueID] = m
	return nil
}

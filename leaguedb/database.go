package leaguedb

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Global_db LeaguesStorage

type LeaguesMongoDB struct {
	DatabaseURL           string
	DatabaseName          string
	LeaguesCollectionName string
}

func (db *LeaguesMongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"leagueid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (db *LeaguesMongoDB) Add(s League) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).Insert(s)
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

func (db *LeaguesMongoDB) Get(keyID string) (League, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	league := League{}
	allWasGood := true

	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).Find(bson.M{"leagueid": keyID}).One(&league)
	if err != nil {
		allWasGood = false
	}

	return league, allWasGood
}

func (db *LeaguesMongoDB) DisplayLeagueName() string {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	str := ""
	//allWasGood := true

	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).Find(nil).Select(bson.M{"name": 1}).One(str)

	return str
}

func (db *LeaguesMongoDB) FindTeam(team string) string {

	return ""
}

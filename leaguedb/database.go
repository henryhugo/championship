package leaguedb

import (
	"encoding/json"
	"fmt"
	"os"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Global_db LeaguesStorage

type LeaguesMongoDB struct {
	DatabaseURL           string
	DatabaseName          string
	LeaguesCollectionName string
}

type Resultat struct {
	res string
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

func (db *LeaguesMongoDB) Add(l League) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).Insert(l)
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

func (db *LeaguesMongoDB) DisplayLeague() string {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	var nameList []League
	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).Find(nil).All(&nameList)

	out, err := json.MarshalIndent(nameList, " ", " ")
	return string(out)
}

func (db *LeaguesMongoDB) RemoveDocument(keyID string) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//allWasGood := true

	err = session.DB(db.DatabaseName).C(db.LeaguesCollectionName).Remove(bson.M{"leagueid": keyID})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		os.Exit(1)
	}

}

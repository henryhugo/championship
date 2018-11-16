package leaguedb

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
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

package leaguedb

import (
	"fmt"

	mgo "gopkg.in/mgo.v2"
)

var Global_db LeaguesStorage

type LeaguesMongoDB struct {
	DatabaseURL            string
	DatabaseName           string
	StudentsCollectionName string
}

func (db *LeaguesMongoDB) Init() {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	index := mgo.Index{
		Key:        []string{"studentid"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err = session.DB(db.DatabaseName).C(db.StudentsCollectionName).EnsureIndex(index)
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

	err = session.DB(db.DatabaseName).C(db.StudentsCollectionName).Insert(s)
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

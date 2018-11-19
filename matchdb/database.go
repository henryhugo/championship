package matchdb

import (
	"encoding/json"
	"fmt"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var Global_db MatchesStorage

type MatchesMongoDB struct {
	DatabaseURL           string
	DatabaseName          string
	MatchesCollectionName string
}

type Resultat struct {
	res string
}

func (db *MatchesMongoDB) Init() {
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

	err = session.DB(db.DatabaseName).C(db.MatchesCollectionName).EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}

func (db *MatchesMongoDB) Add(m MatchesL) error {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	err = session.DB(db.DatabaseName).C(db.MatchesCollectionName).Insert(m)
	if err != nil {
		fmt.Printf("error in Insert(): %v", err.Error())
		return err
	}

	return nil
}

func (db *MatchesMongoDB) Get(keyID string) (MatchesL, bool) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	matchesL := MatchesL{}
	allWasGood := true

	err = session.DB(db.DatabaseName).C(db.MatchesCollectionName).Find(bson.M{"leagueid": keyID}).One(&matchesL)
	if err != nil {
		allWasGood = false
	}

	return matchesL, allWasGood
}

func (db *MatchesMongoDB) DisplayMatches() string {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//allWasGood := true

	var list []MatchesL
	err = session.DB(db.DatabaseName).C(db.MatchesCollectionName).Find(nil).All(&list)
	out, err := json.MarshalIndent(list, " ", " ")
	return string(out)
}

/*func (db *MatchesMongoDB) RemoveDocument(keyID string) {
	session, err := mgo.Dial(db.DatabaseURL)
	if err != nil {
		panic(err)
	}
	defer session.Close()

	//allWasGood := true

	err = session.DB(db.DatabaseName).C(db.MatchesCollectionName).Remove(bson.M{"leagueID": keyID})
	if err != nil {
		fmt.Printf("remove fail %v\n", err)
		os.Exit(1)
	}

}*/

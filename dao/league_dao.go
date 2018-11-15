package dao

import (
	. "championship/models"
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type LeaguesDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "leagues"
)

// Establish a connection to database
func (l *LeaguesDAO) Connect() {
	session, err := mgo.Dial(l.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(l.Database)
}

// Find list of leagues
func (l *LeaguesDAO) FindAll() ([]League, error) {
	var leagues []League
	err := db.C(COLLECTION).Find(bson.M{}).All(&leagues)
	return leagues, err
}

// Find a league by its id
func (l *LeaguesDAO) FindById(id string) (League, error) {
	var league League
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&league)
	return league, err
}

// Insert a league into database
func (l *LeaguesDAO) Insert(league League) error {
	err := db.C(COLLECTION).Insert(&league)
	return err
}

// Delete an existing league
func (l *LeaguesDAO) Delete(league League) error {
	err := db.C(COLLECTION).Remove(&league)
	return err
}

// Update an existing league
func (l *LeaguesDAO) Update(league League) error {
	err := db.C(COLLECTION).UpdateId(league.LeagueID, &league)
	return err
}

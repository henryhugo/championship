package models

import "gopkg.in/mgo.v2/bson"

type League struct {
	CountryID	bson.ObjectId	`bson:"CountryID" json:"CountryID,omitempty"`
	CountryName	string		`bson:"Country" json:"Country,omitempty"`
	LeagueID	bson.ObjectId	`bson:"LeagueID" json:"LeagueID,omitempty"`
	LeagueName	string		`bson:"League" json:"League,omitempty"`
}

/*type Team struct {
	InfoLeague	League	`bson:"InfoLeague" json:"InfoLeague,omitempty"`
	TeamID		bson.ObjectId	`bson:"TeamID" json:"TeamID,omitempty"`
	TeamName	string	`bson:"Team" json:"Team,omitempty"`
}

type Match struct {
	DateMatch 	Time	`bson:"Date" json:"Date,omitempty"`
	FirstTeam	Team	`bson:"FirstTeam" json:"FirstTeam,omitempty"`
	SecondTeam 	Team	`bson:"SecondTeam" json:"SecondTeam,omitempty"`
}*/



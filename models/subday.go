package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type SubdayRecord struct {
	User   string `json:"user"`
	UserID string `json:"userID"`
	Game   string `json:"game"`
}
type SubdayWinnersHistory struct {
	Date    time.Time      `json:"date"`
	Winners []SubdayRecord `json:"winners"`
}
type Subday struct {
	ID             bson.ObjectId
	ChannelID      string                 `json:"channelID"`
	IsActive       bool                   `json:"isActive"`
	SubsOnly       bool                   `json:"subsOnly"`
	Name           string                 `json:"name"`
	Date           time.Time              `json:"date"`
	Votes          []SubdayRecord         `json:"votes"`
	Winners        []SubdayRecord         `json:"winners"`
	WinnersHistory []SubdayWinnersHistory `json:"winnersHistory"`
}

package repositories

import (
	_ "github.com/mattn/go-sqlite3"
)

type EventRepository struct{}

func (self EventRepository) LogEvent(adId string, action string) bool {
	session := getDBSession()
	_, err := session.Exec("INSERT INTO events VALUES (?,?)", adId, action)
	checkErr(err)
	return true
}

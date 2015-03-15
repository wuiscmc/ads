package repositories

import (
	_ "github.com/lib/pq"
)

type EventRepository struct{}

func (self EventRepository) LogEvent(adId int, action string) bool {
	db := GetDBSession()	
	res, err := db.Exec("INSERT INTO events (ad_id, action) VALUES ($1,$2)", adId, action)
	checkErr(err)
  res.LastInsertId()
	return true
}

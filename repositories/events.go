package repositories

import (
	_ "github.com/lib/pq"
	"strconv"
)

type EventRepository struct{}

func (self EventRepository) LogEvent(adId string, action string) bool {
	db := GetDBSession()
	idAd, _ := strconv.Atoi(adId)
	res, err := db.Exec("INSERT INTO events (ad_id, action) VALUES ($1,$2)", idAd, action)
	checkErr(err)
  res.LastInsertId()
	return true
}

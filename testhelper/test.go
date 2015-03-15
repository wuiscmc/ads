package testhelper

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/wuiscmc/ads/repositories"
	"time"
)

var dbSession *sql.DB

func SetTestMode() {
	getDBSession()
}

func getDBSession() *sql.DB {
	if dbSession == nil {
		db, _ := sql.Open("postgres", "dbname=ads_test sslmode=disable")
		repositories.SetDBSession(db)
		dbSession = db
	}
	return dbSession
}

func CreateAd(title string, desc string, prio int, zone int) {
	db := getDBSession()
	db.Exec("INSERT INTO ads (title, description, priority, zone_id) VALUES ($1,$2,$3,$4)", title, desc, prio, zone)
}

func ResetDB() {
	db := getDBSession()
	db.Exec("DELETE FROM impressions; DELETE FROM ads; DELETE FROM events")
}

type Impression struct {
	Id        int
	AdId      int
	Timestamp time.Time
}

func ListImpressions() []Impression {
	db := getDBSession()
	rows, _ := db.Query("SELECT * FROM impressions")
	var impressions []Impression
	for rows.Next() {
		var impression Impression
		rows.Scan(&impression.Id, &impression.AdId, &impression.Timestamp)
		impressions = append(impressions, impression)
	}
	rows.Close()
	return impressions
}

type Event struct {
	AdId   int
	Action string
}

func FetchLastEventTracked() Event {
	db := getDBSession()
	row := db.QueryRow("SELECT * FROM events ORDER BY time DESC LIMIT 1")
	var event Event
	var time time.Time
	var id int
	row.Scan(&id, &event.AdId, &event.Action, &time)
	return event
}

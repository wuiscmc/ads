package testhelper

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wuiscmc/ads/repositories"
	"time"
)

func getDBSession() *sql.DB {
	db, _ := sql.Open("sqlite3", "../ads_test.db")
	return db
}

func ResetDB() {
	db := getDBSession()
	db.Exec("DELETE FROM impressions")
	db.Exec("DELETE FROM ads")
}

func NewAdRepository() *repositories.AdRepository {
	adRepository := repositories.AdRepository{}
	adRepository.SetSession(getDBSession())
	return &adRepository
}

type Impression struct {
	Id        int
	AdId      string
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

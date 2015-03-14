package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wuiscmc/ads/models"
)

type AdRepository struct {
	session *sql.DB
}

func getDBSession() *sql.DB {
	db, err := sql.Open("sqlite3", "./ads.db")
	checkErr(err)
	return db
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (ar *AdRepository) SetSession(db *sql.DB) {
	ar.session = db
}

func NewAdRepository() *AdRepository {
	return &AdRepository{getDBSession()}
}

func (ar AdRepository) FetchAd(zoneId string) models.Ad {
	ad := models.Ad{}
	rows, err := ar.session.Query("SELECT * FROM ads WHERE zoneField = ? ORDER BY priority DESC LIMIT 1", zoneId)
	checkErr(err)
	for rows.Next() {
		rows.Scan(&ad.Id, &ad.Title, &ad.Description, &ad.Priority, &ad.ZoneField)
	}
	return ad
}

func (ac AdRepository) LogImpression(id string) {
	_, err := ac.session.Exec("INSERT INTO impressions (ad_id) values(?)", id)
	checkErr(err)
}

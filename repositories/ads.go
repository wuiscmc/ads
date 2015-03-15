package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"github.com/wuiscmc/ads/models"
)

type AdRepository struct {
	session *sql.DB
}

func (ar *AdRepository) SetSession(db *sql.DB) {
	ar.session = db
}

func NewAdRepository() *AdRepository {
	return &AdRepository{getDBSession()}
}

func (ar AdRepository) FetchAd(zoneId string) models.Ad {
	ad := models.Ad{}
	row := ar.session.QueryRow("SELECT * FROM ads WHERE zoneField = ? ORDER BY priority DESC LIMIT 1", zoneId)
	err := row.Scan(&ad.Id, &ad.Title, &ad.Description, &ad.Priority, &ad.ZoneField)
	checkErr(err)
	return ad
}

func (ar AdRepository) LogImpression(id string) {
	_, err := ar.session.Exec("INSERT INTO impressions (ad_id) values(?)", id)
	checkErr(err)
}

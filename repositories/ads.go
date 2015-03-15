package repositories

import (
	_ "github.com/lib/pq"
	"github.com/wuiscmc/ads/models"
	"database/sql"
)

type AdRepository struct{}

func (ar AdRepository) FetchAd(zoneId int) models.Ad {
	ad := FetchAdFromDB(zoneId)
	medias := FetchAdMediaFromDB(ad.Id)
	ad.SetMediaFiles(medias)
	return ad
}

func (ar AdRepository) LogImpression(adId int) {
	db := GetDBSession()
	_, err := db.Exec("INSERT INTO impressions (ad_id) VALUES ($1)", adId)
	checkErr(err)
}

func FetchAdFromDB(zoneId int) models.Ad {
	db := GetDBSession()
	ad := models.Ad{}
	row := db.QueryRow("SELECT * FROM ads WHERE zone_id = $1 ORDER BY priority DESC LIMIT 1", zoneId)
	err := row.Scan(&ad.Id, &ad.Title, &ad.Description, &ad.Priority, &ad.ZoneId)
	if err == sql.ErrNoRows {
		ad = models.NullAd()
	}
	return ad
}

func FetchAdMediaFromDB(adId int) []models.Media {
	db := GetDBSession()
	rows, err := db.Query("SELECT * FROM ad_media WHERE ad_id = $1", adId)
	checkErr(err)
	 var mediaFiles []models.Media
	for rows.Next() {
		var file models.Media
		var id, aid int
		err = rows.Scan(&id, &aid, &file.Type, &file.Width, &file.Height, &file.Delivery, &file.Url)
		checkErr(err)
		mediaFiles = append(mediaFiles, file)
	}
	rows.Close()
	 return mediaFiles
}

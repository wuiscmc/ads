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

func SeedDB() {
	db := getDBSession()
	defer db.Close()
	db.Exec("INSERT INTO ads values (?,?,?,?,?)", "1", "title1", "description1", 8, 1)
	db.Exec("INSERT INTO ads values (?,?,?,?,?)", "2", "title2", "description2", 9, 1)
	db.Exec("INSERT INTO ads values (?,?,?,?,?)", "3", "title3", "description3", 10, 1)
	db.Exec("INSERT INTO ads values (?,?,?,?,?)", "4", "title4", "description4", 10, 2)
	db.Exec("INSERT INTO ads values (?,?,?,?,?)", "5", "title5", "description5", 9, 3)
	db.Exec("INSERT INTO ads values (?,?,?,?,?)", "6", "title6", "description6", 8, 2)
}

func ResetDB() {
	db := getDBSession()
	defer db.Close()
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
	defer db.Close()
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

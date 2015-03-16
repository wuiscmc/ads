package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
	"os"
)

var dbSession *sql.DB

func SetDBSession(db *sql.DB) {
	dbSession = db
}

func GetDBSession() *sql.DB {
	if dbSession == nil {
		db, err := sql.Open("postgres", getDBParameters())
		checkErr(err)
		dbSession = db
	}
	return dbSession
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getDBParameters() {
	switch os.Getenv("ENV") {
	case "production":
		return os.Getenv("DATABASE_URL")
	case "test":
		return "dbname=ads_test sslmode=disable"
	}

	return "dbname=ads sslmode=disable"
}

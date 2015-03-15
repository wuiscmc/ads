package repositories

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var dbSession *sql.DB

func SetDBSession(db *sql.DB) {
	dbSession = db
}

func GetDBSession() *sql.DB {
	if dbSession == nil {
		db, err := sql.Open("postgres", "dbname=ads sslmode=disable")
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

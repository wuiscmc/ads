package repositories

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

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

package repo

import (
	"database/sql"
	"log"
)

func openDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "vuongnguyen3:vuong@(127.0.0.1:3306)/test_build_web_go")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, err
}

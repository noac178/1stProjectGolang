package repo

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDb() (*sql.DB, error) {
	db, err := sql.Open("mysql", "vuongnguyen3:vuong@(127.0.0.1:3306)/1stpj")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, err
}

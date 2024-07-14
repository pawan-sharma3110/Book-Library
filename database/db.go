package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func DbIN() (db *sql.DB, err error) {
	connStr := `host=localhost port=5432 user=postgres password=Pawan@2003 dbname=library sslmode=disable `
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Panicf("Error in connection string :%v", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		log.Panic(err)
		return nil, err
	}
	return db, nil
}

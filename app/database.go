package app

import (
	"database/sql"
	"fmt"
	"time"
)

func NewDb() *sql.DB {
	dbUser := "root"
	dbPass := ""
	dbName := "journey"
	dbHost := "localhost"
	dbPort := "3306" // Default MySQL port

	// Create a DSN (Data Source Name)
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetConnMaxIdleTime(time.Minute * 10)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db
}

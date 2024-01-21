package app

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *sql.DB {
	config, _ := SetDbConfiguration()

	dsn := fmt.Sprintf("host=%s user=%s port=%s password=%s database=%s sslmode=disable", config.DBHost, config.DBUser, config.DBPort, config.DBPassword, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	sqlDb, err := db.DB()
	if err != nil {
		panic("Failed to get the database object!")
	}
	defer sqlDb.Close()
	return sqlDb
}

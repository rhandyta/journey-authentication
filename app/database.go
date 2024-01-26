package app

import (
	"fmt"
	"journey-user/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDb() *gorm.DB {
	config, _ := SetDbConfiguration()

	dsn := fmt.Sprintf("host=%s user=%s port=%s database=%s sslmode=disable TimeZone=Asia/Shanghai", config.DBHost, config.DBUser, config.DBPort, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	db.AutoMigrate(&model.User{})
	return db
}

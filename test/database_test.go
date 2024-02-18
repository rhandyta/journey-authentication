package test

import (
	"fmt"
	"journey-user/app"
	"journey-user/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestNewDb(t *testing.T) {
	dbSql := NewDbIntegration()
	assert.NotNil(t, dbSql)
}

func NewDbIntegration() *gorm.DB {
	config, err := app.SetDbConfiguration(".env")

	dsn := fmt.Sprintf("host=%s user=%s port=%s database=%s sslmode=disable TimeZone=Asia/Jakarta", config.DBHost, config.DBUser, config.DBPort, config.DBName)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database!")
	}
	db.AutoMigrate(&model.User{})
	return db
}

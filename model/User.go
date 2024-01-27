package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname string         `json:"firstname" Form:"firstname" binding:"required"`
	Lastname  string         `json:"lastname" binding:"required"`
	Age       uint8          `json:"age" binding:"required"`
	Email     string         `json:"email" binding:"required"`
	Password  string         `json:"password" binding:"required"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) TableName() string {
	return "Users"
}

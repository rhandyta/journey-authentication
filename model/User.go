package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname string         `json:"first_name" binding:"required"`
	Lastname  string         `json:"last_name" binding:"required"`
	Age       uint8          `json:"age" binding:"required"`
	Email     string         `gorm:"unique,email" json:"email" binding:"required"`
	Username  string         `gorm:"unique" json:"username" binding:"required"`
	Password  string         `json:"password,omitempty" binding:"required"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) TableName() string {
	return "Users"
}

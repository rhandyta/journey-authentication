package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname string         `json:"firstname"`
	Lastname  string         `json:"lastname"`
	Age       uint8          `json:"age"`
	Email     string         `json:"email"`
	Password  string         `json:"password"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type UserResponse struct {
	Id        uint   `json:"id"`
	Firstname string `json:"firstName"`
	Lastname  string `json:"lastName"`
	Age       uint8  `json:"age"`
	Email     string `json:"email"`
}

func (u *User) TableName() string {
	return "Users"
}

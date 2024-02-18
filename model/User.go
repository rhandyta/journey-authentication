package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	Id        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Firstname string         `json:"first_name" validate:"required"`
	Lastname  string         `json:"last_name" validate:"required"`
	Age       uint8          `json:"age" validate:"required"`
	Email     string         `gorm:"unique,email" json:"email" validate:"required"`
	Username  string         `gorm:"unique" json:"username" validate:"required"`
	Password  string         `json:"password,omitempty" validate:"required"`
	CreatedAt time.Time      `gorm:"<-:create;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"<-;autoCreateTime;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (u *User) TableName() string {
	return "Users"
}

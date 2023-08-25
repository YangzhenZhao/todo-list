package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

func (User) TableName() string {
	return "user"
}

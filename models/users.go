package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string
	Password string
}

type IUserDao interface {
	CreateUser(email, password string) (*User, error)
}

var UserDao IUserDao

type userDaoImpl struct {
	db *gorm.DB
}

func (impl *userDaoImpl) CreateUser(email, password string) (*User, error) {
	_, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func NewUserDao(db *gorm.DB) IUserDao {
	return &userDaoImpl{
		db: db,
	}
}

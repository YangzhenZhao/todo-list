package dao

import (
	"github.com/YangzhenZhao/todo-list/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type IUserDao interface {
	CreateUser(email, password string) error
	GetUserByEmail(email string) (*models.User, error)
}

var UserDao IUserDao

type userDaoImpl struct {
	db *gorm.DB
}

func (impl *userDaoImpl) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	result := impl.db.Where(&models.User{Email: email}).First(&user)
	return user, result.Error
}

func (impl *userDaoImpl) CreateUser(email, password string) error {
	hashPwd, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user := &models.User{
		Email:    email,
		Password: string(hashPwd),
	}
	result := impl.db.Create(user)
	if result.Error != nil || result.RowsAffected == 0 {
		logger.Warn("create user failed, err:%v", result.Error)
		return result.Error
	}
	return nil
}

func NewUserDao(db *gorm.DB) IUserDao {
	return &userDaoImpl{
		db: db,
	}
}

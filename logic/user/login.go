package user

import (
	"github.com/YangzhenZhao/todo-list/dao"
	"github.com/YangzhenZhao/todo-list/dto"
	"golang.org/x/crypto/bcrypt"
)

func Login(loginReq *dto.LoginRequest) (uint, error) {
	user, err := dao.UserDao.GetUserByEmail(loginReq.Email)
	if err != nil {
		return 0, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return 0, err
	}
	return user.ID, nil
}

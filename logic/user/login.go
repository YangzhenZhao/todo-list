package user

import (
	"github.com/YangzhenZhao/todo-list/dao"
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/logic/common"
	"golang.org/x/crypto/bcrypt"
)

func Login(loginReq *dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := dao.UserDao.GetUserByEmail(loginReq.Email)
	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		return nil, err
	}

	token, err := common.GenerateSignedJWT(user.ID)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		UserID: user.ID,
		Token:  token,
	}, nil
}

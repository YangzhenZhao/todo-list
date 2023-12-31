package common

import (
	"time"

	"github.com/YangzhenZhao/todo-list/common/config"
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/golang-jwt/jwt/v5"
)

const tokenTTLInDay = 20 * 24 * time.Hour

func GenerateSignedJWT(userID uint) (string, error) {
	claims := dto.MyCustomClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTLInDay)),
			Issuer:    "project",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.MyTestSigningKey)
}

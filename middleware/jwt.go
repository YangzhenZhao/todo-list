package middleware

import (
	"errors"
	"strings"

	"github.com/YangzhenZhao/todo-list/common/config"
	"github.com/YangzhenZhao/todo-list/common/response"
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthJWTMiddleware(c *gin.Context) {
	headerAuth := c.Request.Header.Get("Authorization")
	if headerAuth == "" {
		response.UnauthorizedResponse(c, "no auth in header")
		c.Abort()
		return
	}

	splitBearerAuths := strings.SplitN(headerAuth, " ", 2)
	if len(splitBearerAuths) != 2 || splitBearerAuths[0] != "Bearer" {
		response.UnauthorizedResponse(c, "invalid auth format")
		c.Abort()
		return
	}

	claims, err := parseToken(splitBearerAuths[1])
	if err != nil {
		response.UnauthorizedResponse(c, "invalid token")
		c.Abort()
		return
	}

	c.Set("userID", claims.UserID)
}

func parseToken(tokenString string) (*dto.MyCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &dto.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return config.MyTestSigningKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*dto.MyCustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

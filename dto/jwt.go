package dto

import "github.com/golang-jwt/jwt/v5"

type MyCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

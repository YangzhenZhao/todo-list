package dto

import "encoding/json"

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	UserID uint   `json:"userID"`
	Token  string `json:"token"`
}

func (res *LoginResponse) JsonDumps() string {
	dumpsRes, _ := json.Marshal(res)
	return string(dumpsRes)
}

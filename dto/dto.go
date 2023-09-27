package dto

import (
	"encoding/json"
	"time"
)

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

type CreateTodoRequest struct {
	UserID      uint      `json:"userID"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	IsFinished  bool      `json:"isFinished"`
	IsStar      bool      `json:"isStar"`
	Priority    string    `json:"priority"`
	Location    string    `json:"location"`
}

type CreateTodoResponse struct {
	TodoID uint `json:"todoID"`
}

func (res *CreateTodoResponse) JsonDumps() string {
	dumpsRes, _ := json.Marshal(res)
	return string(dumpsRes)
}

type DeleteTodoRequest struct {
	UserID uint `json:"userID"`
	TodoID uint `json:"todoID"`
}

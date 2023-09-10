package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
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

func (Todo) TableName() string {
	return "todo"
}

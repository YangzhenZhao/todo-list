package models

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	UserID        uint      `json:"userID"`
	Title         string    `json:"title"`
	Description   string    `json:"description"`
	StartDatetime time.Time `json:"start_datetime"`
	EndDatetime   time.Time `json:"end_datetime"`
	IsFinished    bool      `json:"isFinished"`
	IsStar        bool      `json:"isStar"`
	Priority      string    `json:"priority"`
	Location      string    `json:"location"`
}

func (Todo) TableName() string {
	return "todo"
}

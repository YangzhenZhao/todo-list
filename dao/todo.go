package dao

import (
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/models"
	"gorm.io/gorm"
)

type ITodoDao interface {
	CreateTodo(todo *dto.CreateTodoRequest) error
}

var TodoDao ITodoDao

type todoDaoImpl struct {
	db *gorm.DB
}

func (impl *todoDaoImpl) CreateTodo(todo *dto.CreateTodoRequest) error {
	todoModel := &models.Todo{
		UserID:        todo.UserID,
		Title:         todo.Title,
		Description:   todo.Description,
		StartDatetime: todo.StartTime,
		EndDatetime:   todo.EndTime,
		IsFinished:    todo.IsFinished,
		IsStar:        todo.IsStar,
		Priority:      todo.Priority,
		Location:      todo.Location,
	}
	result := impl.db.Create(todoModel)
	if result.Error != nil || result.RowsAffected == 0 {
		logger.Warn("create todo failed, err:%v", result.Error)
		return result.Error
	}
	return nil
}

func NewTodoDao(db *gorm.DB) ITodoDao {
	return &todoDaoImpl{
		db: db,
	}
}

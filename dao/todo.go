package dao

import (
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/models"
	"gorm.io/gorm"
)

type ITodoDao interface {
	CreateTodo(todo *dto.CreateTodoRequest) (uint, error)
}

var TodoDao ITodoDao

type todoDaoImpl struct {
	db *gorm.DB
}

func (impl *todoDaoImpl) CreateTodo(todo *dto.CreateTodoRequest) (uint, error) {
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
		logger.Warn("create todo failed, err", result.Error)
		return 0, result.Error
	}
	logger.Info("[dao CreateTodo] todoID:", todoModel.ID)
	return todoModel.ID, nil
}

func NewTodoDao(db *gorm.DB) ITodoDao {
	return &todoDaoImpl{
		db: db,
	}
}

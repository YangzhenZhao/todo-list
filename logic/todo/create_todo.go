package todo

import (
	"github.com/YangzhenZhao/todo-list/dao"
	"github.com/YangzhenZhao/todo-list/dto"
)

func CreateTodo(todo *dto.CreateTodoRequest) (uint, error) {
	return dao.TodoDao.CreateTodo(todo)
}

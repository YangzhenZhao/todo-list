package todo

import "github.com/YangzhenZhao/todo-list/dao"

func DeleteTodo(todoID uint) error {
	return dao.TodoDao.DeleteTodo(todoID)
}

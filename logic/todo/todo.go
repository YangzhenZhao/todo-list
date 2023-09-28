package todo

import "github.com/YangzhenZhao/todo-list/dao"

func Star(todoID uint) error {
	return dao.TodoDao.Star(todoID)
}

func Unstar(todoID uint) error {
	return dao.TodoDao.Unstar(todoID)
}

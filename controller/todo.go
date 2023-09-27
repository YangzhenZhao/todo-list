package controller

import (
	"github.com/YangzhenZhao/todo-list/common/response"
	"github.com/YangzhenZhao/todo-list/dto"
	"github.com/YangzhenZhao/todo-list/logic/todo"
	"github.com/davecgh/go-spew/spew"
	"github.com/gin-gonic/gin"
)

func GetTodolist(c *gin.Context) {

}

func CreateTodo(c *gin.Context) {
	createReq := &dto.CreateTodoRequest{}
	if err := unmarshalRequest(c, createReq); err != nil {
		logger.Info("createTodo request, err:", err)
		response.InvalidArgumentResponse(c, "参数不合法")
		return
	}

	if !validteToken(c, createReq.UserID, "CreateTodo") {
		return
	}

	logger.Info("[CreateTodo] req", spew.Sdump(createReq))

	todoID, err := todo.CreateTodo(createReq)
	if err != nil {
		response.InternalServerErrorResponse(c, err.Error())
		logger.Info("createTodo internal server error", err)
		return
	}

	createTodoRes := &dto.CreateTodoResponse{
		TodoID: todoID,
	}
	response.SuccessResponse(c, createTodoRes.JsonDumps(), "创建成功")
}

func UpdateTodo(c *gin.Context) {

}

func DeleteTodo(c *gin.Context) {
	delReq := &dto.DeleteTodoRequest{}
	if err := unmarshalRequest(c, delReq); err != nil {
		logger.Info("DeleteTodo request, err:", err)
		response.InvalidArgumentResponse(c, "参数不合法")
		return
	}

	if !validteToken(c, delReq.UserID, "DeleteTodo") {
		return
	}

	logger.Info("[DeleteTodo] req:", spew.Sdump(delReq))

	err := todo.DeleteTodo(delReq.TodoID)
	if err != nil {
		response.InternalServerErrorResponse(c, err.Error())
		logger.Info("deleteTodo internal server error", err)
		return
	}

	response.SuccessResponse(c, "", "删除成功")
}

func GetTodo(c *gin.Context) {

}

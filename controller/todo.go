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
	funcName := "CreateTodo"

	createReq := &dto.CreateTodoRequest{}
	if !validteReq(c, createReq, funcName) {
		return
	}
	if !validteToken(c, createReq.UserID, funcName) {
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
	funcName := "DeleteTodo"

	delReq := &dto.DeleteTodoRequest{}
	if !validteReq(c, delReq, funcName) {
		return
	}
	if !validteToken(c, delReq.UserID, funcName) {
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

func StarTodo(c *gin.Context) {
	funcName := "StarTodo"

	starReq := &dto.SetStarRequest{}
	if !validteReq(c, starReq, funcName) {
		return
	}
	if !validteToken(c, starReq.UserID, funcName) {
		return
	}

	err := todo.Star(starReq.TodoID)
	if err != nil {
		response.InternalServerErrorResponse(c, err.Error())
		logger.Info(funcName, " internal server error", err)
		return
	}

	response.SuccessResponse(c, "", "star成功")
}

func UnstarTodo(c *gin.Context) {
	funcName := "UnstarTodo"

	unstarReq := &dto.SetStarRequest{}
	if !validteReq(c, unstarReq, funcName) {
		return
	}
	if !validteToken(c, unstarReq.UserID, funcName) {
		return
	}

	err := todo.Unstar(unstarReq.TodoID)
	if err != nil {
		response.InternalServerErrorResponse(c, err.Error())
		logger.Info(funcName, " internal server error", err)
		return
	}

	response.SuccessResponse(c, "", "unstar成功")
}

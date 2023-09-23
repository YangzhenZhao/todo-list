package controller

import (
	"log"

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
		log.Printf("createTodo request err: %v\n", err)
		response.InvalidArgumentResponse(c, "参数不合法")
		return
	}
	tokenUserID, exists := c.Get("userID")
	if !exists {
		log.Printf("createTodo token don't have userID\n")
		response.InvalidArgumentResponse(c, "token不合法")
		return
	}
	if createReq.UserID != tokenUserID {
		log.Printf("createTodo invalid userID, createReq.UserID:%d, tokenUserID:%d\n", createReq.UserID, tokenUserID)
		response.InvalidArgumentResponse(c, "token不合法")
		return
	}

	log.Printf("[CreateTodo] req: %s\n", spew.Sdump(createReq))

	err := todo.CreateTodo(createReq)
	if err != nil {
		response.InternalServerErrorResponse(c, err.Error())
		log.Printf("createTodo internal server error: %v\n", err)
		return
	}

	response.SuccessResponse(c, "", "创建成功")
}

func UpdateTodo(c *gin.Context) {

}

func DeleteTodo(c *gin.Context) {

}

func GetTodo(c *gin.Context) {

}

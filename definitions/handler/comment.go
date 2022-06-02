package handler_i

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/request"
)

type Comment interface {
	List(ctx *gin.Context) (request.CommentList, error)
}

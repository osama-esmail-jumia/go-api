package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/request"
)

type CommentI interface {
	List(ctx *gin.Context) (request.CommentList, error)
}

type Comment struct {
}

func NewComment() Comment {
	return Comment{}
}

func (h Comment) List(ctx *gin.Context) (request.CommentList, error) {
	var req request.CommentList
	err := ctx.ShouldBind(&req)
	return req, err
}

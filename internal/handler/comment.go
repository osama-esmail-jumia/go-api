package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/http_model"
)

type Comment struct {
}

func NewComment() Comment {
	return Comment{}
}

func (h Comment) List(ctx *gin.Context) (http_model.CommentListRequest, error) {
	var req http_model.CommentListRequest
	err := ctx.ShouldBind(&req)
	return req, err
}

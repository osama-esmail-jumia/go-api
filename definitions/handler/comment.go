package handler_i

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/http_model"
)

type Comment interface {
	List(ctx *gin.Context) (http_model.CommentListRequest, error)
}

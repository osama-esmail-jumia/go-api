package comment

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/comment"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) List(ctx *gin.Context) (comment_i.ListRequest, error) {
	var req comment_i.ListRequest
	err := ctx.ShouldBind(&req)
	return req, err
}

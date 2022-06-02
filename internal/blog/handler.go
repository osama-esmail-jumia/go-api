package blog

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/blog"
)

type Handler struct {
}

func NewHandler() Handler {
	return Handler{}
}

func (h Handler) Create(ctx *gin.Context) (blog_i.CreateRequest, error) {
	var req blog_i.CreateRequest
	err := ctx.ShouldBindJSON(&req)
	return req, err
}

func (h Handler) Update(ctx *gin.Context) (req blog_i.UpdateRequest, err error) {
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		return
	}
	err = ctx.ShouldBindJSON(&req)
	return
}

func (h Handler) Delete(ctx *gin.Context) (blog_i.DeleteRequest, error) {
	var req blog_i.DeleteRequest
	err := ctx.ShouldBindUri(&req)
	return req, err
}

func (h Handler) List(ctx *gin.Context) (blog_i.ListRequest, error) {
	var req blog_i.ListRequest
	err := ctx.ShouldBind(&req)
	return req, err
}

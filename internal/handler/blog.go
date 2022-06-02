package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/http_model"
)

type Blog struct {
}

func NewBlog() Blog {
	return Blog{}
}

func (h Blog) Create(ctx *gin.Context) (http_model.BlogCreateRequest, error) {
	var req http_model.BlogCreateRequest
	err := ctx.ShouldBindJSON(&req)
	return req, err
}

func (h Blog) Update(ctx *gin.Context) (req http_model.BlogUpdateRequest, err error) {
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		return
	}
	err = ctx.ShouldBindJSON(&req)
	return
}

func (h Blog) Delete(ctx *gin.Context) (http_model.BlogDeleteRequest, error) {
	var req http_model.BlogDeleteRequest
	err := ctx.ShouldBindUri(&req)
	return req, err
}

func (h Blog) List(ctx *gin.Context) (http_model.BlogListRequest, error) {
	var req http_model.BlogListRequest
	err := ctx.ShouldBind(&req)
	return req, err
}

package handler

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/request"
)

type Blog struct {
}

func NewBlog() Blog {
	return Blog{}
}

func (h Blog) Create(ctx *gin.Context) (request.BlogCreate, error) {
	var req request.BlogCreate
	err := ctx.ShouldBindJSON(&req)
	return req, err
}

func (h Blog) Update(ctx *gin.Context) (req request.BlogUpdate, err error) {
	err = ctx.ShouldBindUri(&req)
	if err != nil {
		return
	}
	err = ctx.ShouldBindJSON(&req)
	return
}

func (h Blog) Delete(ctx *gin.Context) (request.BlogDelete, error) {
	var req request.BlogDelete
	err := ctx.ShouldBindUri(&req)
	return req, err
}

func (h Blog) List(ctx *gin.Context) (request.BlogList, error) {
	var req request.BlogList
	err := ctx.ShouldBind(&req)
	return req, err
}

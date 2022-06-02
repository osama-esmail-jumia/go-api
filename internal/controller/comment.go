package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/internal/handler"
	"go-api/internal/service"
	"go-api/pkg/api"
	"net/http"
)

type CommentI interface {
	List(ctx *gin.Context)
}

type Comment struct {
	handler handler.CommentI
	service service.CommentI
}

func NewComment(handler handler.CommentI, service service.CommentI) Comment {
	return Comment{handler, service}
}

func (c Comment) List(ctx *gin.Context) {
	req, err := c.handler.List(ctx)
	
	if err != nil {
		fmt.Println(err)
		api.HandleResponseError(ctx, err)
		return
	}
	
	resp, err := c.service.List(req)
	
	if err != nil {
		fmt.Println(err)
		api.HandleResponseError(ctx, err)
		return
	}
	
	api.HandleResponse(
		ctx,
		http.StatusCreated,
		api.Response{Data: resp},
	)
}

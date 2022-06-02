package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/definitions/handler"
	"go-api/definitions/service"
	"go-api/pkg/api"
	"net/http"
)

type Comment struct {
	handler handler_i.Comment
	service service_i.Comment
}

func NewComment(handler handler_i.Comment, service service_i.Comment) Comment {
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

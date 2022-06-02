package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/definitions/handler"
	"go-api/definitions/service"
	"go-api/pkg/api"
	"net/http"
)

type Blog struct {
	handler handler_i.Blog
	service service_i.Blog
}

func NewBlog(handler handler_i.Blog, service service_i.Blog) Blog {
	return Blog{handler, service}
}

func (c Blog) Create(ctx *gin.Context) {
	req, err := c.handler.Create(ctx)
	
	if err != nil {
		fmt.Println(err)
		api.HandleResponseError(ctx, err)
		return
	}
	
	resp, err := c.service.Create(req)
	
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

func (c Blog) Update(ctx *gin.Context) {
	req, err := c.handler.Update(ctx)
	
	if err != nil {
		fmt.Println(err)
		api.HandleResponseError(ctx, err)
		return
	}
	
	resp, err := c.service.Update(req)
	
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

func (c Blog) Delete(ctx *gin.Context) {
	req, err := c.handler.Delete(ctx)
	
	if err != nil {
		fmt.Println(err)
		api.HandleResponseError(ctx, err)
		return
	}
	
	err = c.service.Delete(req)
	
	if err != nil {
		fmt.Println(err)
		api.HandleResponseError(ctx, err)
		return
	}
	
	api.HandleResponse(
		ctx,
		http.StatusCreated,
		api.Response{},
	)
}

func (c Blog) List(ctx *gin.Context) {
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

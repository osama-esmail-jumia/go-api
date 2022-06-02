package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/internal/handler"
	"go-api/internal/service"
	"go-api/pkg/api"
	"net/http"
)

type BlogI interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

type Blog struct {
	handler handler.BlogI
	service service.BlogI
}

func NewBlog(handler handler.BlogI, service service.BlogI) Blog {
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

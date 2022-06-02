package blog

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/definitions/blog"
	"go-api/pkg/api"
	"net/http"
)

type Controller struct {
	handler blog_i.Handler
	service blog_i.Service
}

func NewController(handler blog_i.Handler, service blog_i.Service) Controller {
	return Controller{handler, service}
}

func (c Controller) Create(ctx *gin.Context) {
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

func (c Controller) Update(ctx *gin.Context) {
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

func (c Controller) Delete(ctx *gin.Context) {
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

func (c Controller) List(ctx *gin.Context) {
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

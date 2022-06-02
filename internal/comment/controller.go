package comment

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-api/definitions/comment"
	"go-api/pkg/api"
	"net/http"
)

type Controller struct {
	handler comment_i.Handler
	service comment_i.Service
}

func NewController(handler comment_i.Handler, service comment_i.Service) Controller {
	return Controller{handler, service}
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

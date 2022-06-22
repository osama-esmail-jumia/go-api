package router

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/controller"
)

func NewComment(rg *gin.RouterGroup, ctrl controller.Comment) {
	r := rg.Group("comment")
	
	r.GET("", ctrl.List)
}

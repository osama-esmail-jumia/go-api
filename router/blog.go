package router

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/controller"
)

func NewBlog(rg *gin.RouterGroup, ctrl controller.BlogI) {
	r := rg.Group("blog")
	
	r.POST("", ctrl.Create)
	r.PATCH(":id", ctrl.Update)
	r.DELETE(":id", ctrl.Delete)
	r.GET("", ctrl.List)
}

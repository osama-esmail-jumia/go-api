package router

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/blog"
)

func NewBlog(rg *gin.RouterGroup, ctrl blog_i.Controller) {
	r := rg.Group("blog")
	
	r.POST("", ctrl.Create)
	r.PATCH(":id", ctrl.Update)
	r.DELETE(":id", ctrl.Delete)
	r.GET("", ctrl.List)
}

package router

import (
	"github.com/gin-gonic/gin"
	"go-api/internal/controller"
)

func NewComment(rg *gin.RouterGroup, ctrl controller.CommentI) {
	r := rg.Group("comment")
	
	r.GET("", ctrl.List)
}

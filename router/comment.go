package router

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/comment"
)

func NewComment(rg *gin.RouterGroup, ctrl comment_i.Controller) {
	r := rg.Group("comment")
	
	r.GET("", ctrl.List)
}

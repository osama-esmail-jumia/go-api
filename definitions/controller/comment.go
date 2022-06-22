package controller

import (
	"github.com/gin-gonic/gin"
)

type Comment interface {
	List(ctx *gin.Context)
}

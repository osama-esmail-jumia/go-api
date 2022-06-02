package controller_i

import (
	"github.com/gin-gonic/gin"
)

type Blog interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

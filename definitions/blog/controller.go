package blog_i

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	List(ctx *gin.Context)
}

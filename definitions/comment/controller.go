package comment_i

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	List(ctx *gin.Context)
}

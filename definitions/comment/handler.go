package comment_i

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	List(ctx *gin.Context) (ListRequest, error)
}

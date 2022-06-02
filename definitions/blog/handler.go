package blog_i

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	Create(ctx *gin.Context) (CreateRequest, error)
	Update(ctx *gin.Context) (UpdateRequest, error)
	Delete(ctx *gin.Context) (DeleteRequest, error)
	List(ctx *gin.Context) (ListRequest, error)
}

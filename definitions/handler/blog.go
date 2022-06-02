package handler_i

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/http_model"
)

type Blog interface {
	Create(ctx *gin.Context) (http_model.BlogCreateRequest, error)
	Update(ctx *gin.Context) (http_model.BlogUpdateRequest, error)
	Delete(ctx *gin.Context) (http_model.BlogDeleteRequest, error)
	List(ctx *gin.Context) (http_model.BlogListRequest, error)
}

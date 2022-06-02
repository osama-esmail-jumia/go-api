package handler_i

import (
	"github.com/gin-gonic/gin"
	"go-api/definitions/request"
)

type Blog interface {
	Create(ctx *gin.Context) (request.BlogCreate, error)
	Update(ctx *gin.Context) (request.BlogUpdate, error)
	Delete(ctx *gin.Context) (request.BlogDelete, error)
	List(ctx *gin.Context) (request.BlogList, error)
}

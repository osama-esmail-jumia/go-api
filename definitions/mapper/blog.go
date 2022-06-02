package mapper_i

import (
	"go-api/definitions/model"
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Blog interface {
	CreateRequestToModel(req request.BlogCreate) model.Blog
	UpdateRequestToModel(req request.BlogUpdate) model.Blog
	DeleteRequestToModel(req request.BlogDelete) model.Blog
	ModelToResponse(model model.Blog) response.Blog
	ModelToListResponse(model []model.Blog) response.BlogList
	ListRequestToFilter(req request.BlogList) model.BlogFilter
}

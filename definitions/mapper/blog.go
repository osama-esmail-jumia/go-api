package mapper_i

import (
	"go-api/definitions/http_model"
	"go-api/definitions/model"
)

type Blog interface {
	CreateRequestToModel(req http_model.BlogCreateRequest) model.Blog
	UpdateRequestToModel(req http_model.BlogUpdateRequest) model.Blog
	DeleteRequestToModel(req http_model.BlogDeleteRequest) model.Blog
	ModelToResponse(model model.Blog) http_model.BlogResponse
	ModelToListResponse(model []model.Blog) http_model.BlogListResponse
	ListRequestToFilter(req http_model.BlogListRequest) model.BlogFilter
}

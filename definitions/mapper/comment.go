package mapper_i

import (
	"go-api/definitions/http_model"
	"go-api/definitions/model"
)

type Comment interface {
	ModelToResponse(model model.Comment) http_model.CommentResponse
	ModelToListResponse(model []model.Comment) http_model.CommentListResponse
	ListRequestToFilter(req http_model.CommentListRequest) model.CommentFilter
}

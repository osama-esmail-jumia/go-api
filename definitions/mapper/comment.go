package mapper_i

import (
	"go-api/definitions/model"
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Comment interface {
	ModelToResponse(model model.Comment) response.Comment
	ModelToListResponse(model []model.Comment) response.CommentList
	ListRequestToFilter(req request.CommentList) model.CommentFilter
}

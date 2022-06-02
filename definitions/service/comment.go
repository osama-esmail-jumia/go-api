package service_i

import (
	"go-api/definitions/http_model"
)

type Comment interface {
	List(req http_model.CommentListRequest) (http_model.CommentListResponse, error)
}

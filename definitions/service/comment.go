package service_i

import (
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Comment interface {
	List(req request.CommentList) (response.CommentList, error)
}

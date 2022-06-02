package repository_i

import "go-api/definitions/model"

type Comment interface {
	List(models *[]model.Comment, filter model.CommentFilter) error
}

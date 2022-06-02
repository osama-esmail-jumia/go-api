package blog_i

import "go-api/definitions/comment"

type Model struct {
	ID          int64
	Title       *string
	Description *string
	Comments    []comment_i.Model
	target      string `gorm:"-"`
}

type ModelFilter struct {
	Title       *string
	Description *string
}

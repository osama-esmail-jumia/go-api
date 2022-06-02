package comment_i

import "go-api/definitions/blog"

type Model struct {
	ID     int64
	BlogID int64
	Text   *string
	Blog   *blog_i.Model
}

type ModelFilter struct {
	BlogID *int64
	Text   *string
}

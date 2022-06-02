package blog_i

import "go-api/definitions/comment"

type Response struct {
	ID          int64                `json:"id"`
	Title       string               `json:"title"`
	Description string               `json:"description"`
	Comments    []comment_i.Response `json:"comments"`
}

type ListResponse []Response

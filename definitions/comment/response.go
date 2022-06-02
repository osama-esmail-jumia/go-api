package comment_i

import "go-api/definitions/blog"

type Response struct {
	ID     int64            `json:"id"`
	BlogID int64            `json:"blog_id"`
	Text   string           `json:"text"`
	Blog   *blog_i.Response `json:"blog,omitempty"`
}

type ListResponse []Response

package http_model

type CommentListRequest struct {
	BlogID *int64  `form:"blog_id"`
	Text   *string `form:"text"`
}

type CommentResponse struct {
	ID     int64         `json:"id"`
	BlogID int64         `json:"blog_id"`
	Text   string        `json:"text"`
	Blog   *BlogResponse `json:"blog,omitempty"`
}

type CommentListResponse []CommentResponse

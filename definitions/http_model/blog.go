package http_model

type BlogCreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type BlogUpdateRequest struct {
	ID          int64   `uri:"id" binding:"required"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type BlogDeleteRequest struct {
	ID int64 `uri:"id" form:"required"`
}

type BlogListRequest struct {
	Title       *string `form:"title"`
	Description *string `form:"description"`
}

type BlogResponse struct {
	ID          int64             `json:"id"`
	Title       string            `json:"title"`
	Description string            `json:"description"`
	Comments    []CommentResponse `json:"comments"`
}

type BlogListResponse []BlogResponse

package blog_i

type CreateRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type UpdateRequest struct {
	ID          int64   `uri:"id" binding:"required"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type DeleteRequest struct {
	ID int64 `uri:"id" form:"required"`
}

type ListRequest struct {
	Title       *string `form:"title"`
	Description *string `form:"description"`
}

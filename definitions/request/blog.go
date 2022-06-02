package request

type BlogCreate struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
}

type BlogUpdate struct {
	ID          int64   `uri:"id" binding:"required"`
	Title       *string `json:"title"`
	Description *string `json:"description"`
}

type BlogDelete struct {
	ID int64 `uri:"id" form:"required"`
}

type BlogList struct {
	Title       *string `form:"title"`
	Description *string `form:"description"`
}

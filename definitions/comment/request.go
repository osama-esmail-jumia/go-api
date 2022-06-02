package comment_i

type ListRequest struct {
	BlogID *int64  `form:"blog_id"`
	Text   *string `form:"text"`
}

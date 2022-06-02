package request

type CommentList struct {
	BlogID *int64  `form:"blog_id"`
	Text   *string `form:"text"`
}

package response

type Comment struct {
	ID     int64  `json:"id"`
	BlogID int64  `json:"blog_id"`
	Text   string `json:"text"`
	Blog   *Blog  `json:"blog,omitempty"`
}

type CommentList []Comment

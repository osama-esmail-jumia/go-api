package model

type Comment struct {
	ID     int64
	BlogID int64
	Text   *string
	Blog   *Blog
}

type CommentFilter struct {
	BlogID *int64
	Text   *string
}

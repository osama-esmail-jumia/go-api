package comment_i

type Model struct {
	ID     int64
	BlogID int64
	Text   *string
	// Circular dependency error
	//Blog   *blog_i.Model
}

type ModelFilter struct {
	BlogID *int64
	Text   *string
}

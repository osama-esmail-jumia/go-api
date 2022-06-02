package comment_i

type Response struct {
	ID     int64  `json:"id"`
	BlogID int64  `json:"blog_id"`
	Text   string `json:"text"`
	// Circular dependency error
	//Blog   *blog_i.Response `json:"blog,omitempty"`
}

type ListResponse []Response

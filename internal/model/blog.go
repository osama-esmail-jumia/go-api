package model

type Blog struct {
	ID          int64
	Title       *string
	Description *string
	Comments    []Comment
	target      string `gorm:"-"`
}

type BlogFilter struct {
	Title       *string
	Description *string
}

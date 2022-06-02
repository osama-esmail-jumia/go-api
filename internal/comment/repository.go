package comment

import (
	comment_i "go-api/definitions/comment"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db}
}

func (r Repository) List(models *[]comment_i.Model, filter comment_i.ModelFilter) error {
	return r.db.Where(filter).Preload("Blog", func() {}).Find(models).Error
}

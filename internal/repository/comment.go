package repository

import (
	"go-api/internal/model"
	"gorm.io/gorm"
)

type CommentI interface {
	List(models *[]model.Comment, filter model.CommentFilter) error
}

type Comment struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) Comment {
	return Comment{db}
}

func (r Comment) List(models *[]model.Comment, filter model.CommentFilter) error {
	return r.db.Where(filter).Preload("Blog", func() {}).Find(models).Error
}

package repository

import (
	"go-api/internal/model"
	"gorm.io/gorm"
)

type BlogI interface {
	Create(model *model.Blog) error
	Delete(model *model.Blog) error
	Update(model *model.Blog) error
	List(models *[]model.Blog, filter model.BlogFilter) error
}

type Blog struct {
	db *gorm.DB
}

func NewBlog(db *gorm.DB) Blog {
	return Blog{db}
}

func (r Blog) Create(model *model.Blog) error {
	return r.db.Create(&model).Error
}

func (r Blog) Update(model *model.Blog) error {
	return r.db.Updates(&model).Error
}

func (r Blog) Delete(model *model.Blog) error {
	return r.db.Delete(model).Error
}

func (r Blog) List(models *[]model.Blog, filter model.BlogFilter) error {
	return r.db.Where(filter).Preload("Comments").Find(models).Error
}

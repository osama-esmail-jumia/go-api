package blog

import (
	"go-api/definitions/blog"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return Repository{db}
}

func (r Repository) Create(model *blog_i.Model) error {
	return r.db.Create(&model).Error
}

func (r Repository) Update(model *blog_i.Model) error {
	return r.db.Updates(&model).Error
}

func (r Repository) Delete(model *blog_i.Model) error {
	return r.db.Delete(model).Error
}

func (r Repository) List(models *[]blog_i.Model, filter blog_i.ModelFilter) error {
	return r.db.Where(filter).Preload("Comments").Find(models).Error
}

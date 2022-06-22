package repository

import "go-api/definitions/model"

type Blog interface {
	Create(model *model.Blog) error
	Delete(model *model.Blog) error
	Update(model *model.Blog) error
	List(models *[]model.Blog, filter model.BlogFilter) error
}

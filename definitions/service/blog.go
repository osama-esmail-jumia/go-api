package service

import (
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Blog interface {
	Create(create request.BlogCreate) (response.Blog, error)
	Update(req request.BlogUpdate) (response.Blog, error)
	Delete(req request.BlogDelete) error
	List(req request.BlogList) (response.BlogList, error)
}

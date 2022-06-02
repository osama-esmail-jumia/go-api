package service_i

import (
	"go-api/definitions/http_model"
)

type Blog interface {
	Create(create http_model.BlogCreateRequest) (http_model.BlogResponse, error)
	Update(req http_model.BlogUpdateRequest) (http_model.BlogResponse, error)
	Delete(req http_model.BlogDeleteRequest) error
	List(req http_model.BlogListRequest) (http_model.BlogListResponse, error)
}

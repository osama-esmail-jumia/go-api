package service

import (
	"go-api/definitions/http_model"
	"go-api/definitions/mapper"
	"go-api/definitions/model"
	"go-api/definitions/repository"
)

type Blog struct {
	repo   repository_i.Blog
	mapper mapper_i.Blog
}

func NewBlog(repo repository_i.Blog, mapper mapper_i.Blog) Blog {
	return Blog{repo, mapper}
}

func (s Blog) Create(req http_model.BlogCreateRequest) (http_model.BlogResponse, error) {
	row := s.mapper.CreateRequestToModel(req)
	err := s.repo.Create(&row)
	resp := s.mapper.ModelToResponse(row)
	
	return resp, err
}

func (s Blog) Update(req http_model.BlogUpdateRequest) (http_model.BlogResponse, error) {
	row := s.mapper.UpdateRequestToModel(req)
	err := s.repo.Update(&row)
	resp := s.mapper.ModelToResponse(row)
	
	return resp, err
}

func (s Blog) Delete(req http_model.BlogDeleteRequest) error {
	row := s.mapper.DeleteRequestToModel(req)
	err := s.repo.Delete(&row)
	s.mapper.ModelToResponse(row)
	
	return err
}

func (s Blog) List(req http_model.BlogListRequest) (http_model.BlogListResponse, error) {
	var rows []model.Blog
	filter := s.mapper.ListRequestToFilter(req)
	err := s.repo.List(&rows, filter)
	resp := s.mapper.ModelToListResponse(rows)
	
	return resp, err
}

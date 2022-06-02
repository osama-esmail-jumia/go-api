package service

import (
	"go-api/definitions/mapper"
	"go-api/definitions/model"
	"go-api/definitions/repository"
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Blog struct {
	repo   repository_i.Blog
	mapper mapper_i.Blog
}

func NewBlog(repo repository_i.Blog, mapper mapper_i.Blog) Blog {
	return Blog{repo, mapper}
}

func (s Blog) Create(req request.BlogCreate) (response.Blog, error) {
	row := s.mapper.CreateRequestToModel(req)
	err := s.repo.Create(&row)
	resp := s.mapper.ModelToResponse(row)
	
	return resp, err
}

func (s Blog) Update(req request.BlogUpdate) (response.Blog, error) {
	row := s.mapper.UpdateRequestToModel(req)
	err := s.repo.Update(&row)
	resp := s.mapper.ModelToResponse(row)
	
	return resp, err
}

func (s Blog) Delete(req request.BlogDelete) error {
	row := s.mapper.DeleteRequestToModel(req)
	err := s.repo.Delete(&row)
	s.mapper.ModelToResponse(row)
	
	return err
}

func (s Blog) List(req request.BlogList) (response.BlogList, error) {
	var rows []model.Blog
	filter := s.mapper.ListRequestToFilter(req)
	err := s.repo.List(&rows, filter)
	resp := s.mapper.ModelToListResponse(rows)
	
	return resp, err
}

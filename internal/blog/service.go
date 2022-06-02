package blog

import (
	"go-api/definitions/blog"
)

type Service struct {
	repo   blog_i.Repository
	mapper blog_i.Mapper
}

func NewService(repo blog_i.Repository, mapper blog_i.Mapper) Service {
	return Service{repo, mapper}
}

func (s Service) Create(req blog_i.CreateRequest) (blog_i.Response, error) {
	row := s.mapper.CreateRequestToModel(req)
	err := s.repo.Create(&row)
	resp := s.mapper.ModelToResponse(row)
	
	return resp, err
}

func (s Service) Update(req blog_i.UpdateRequest) (blog_i.Response, error) {
	row := s.mapper.UpdateRequestToModel(req)
	err := s.repo.Update(&row)
	resp := s.mapper.ModelToResponse(row)
	
	return resp, err
}

func (s Service) Delete(req blog_i.DeleteRequest) error {
	row := s.mapper.DeleteRequestToModel(req)
	err := s.repo.Delete(&row)
	s.mapper.ModelToResponse(row)
	
	return err
}

func (s Service) List(req blog_i.ListRequest) (blog_i.ListResponse, error) {
	var rows []blog_i.Model
	filter := s.mapper.ListRequestToFilter(req)
	err := s.repo.List(&rows, filter)
	resp := s.mapper.ModelToListResponse(rows)
	
	return resp, err
}

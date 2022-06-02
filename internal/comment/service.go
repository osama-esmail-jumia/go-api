package comment

import (
	"go-api/definitions/comment"
)

type Service struct {
	repo   comment_i.Repository
	mapper comment_i.Mapper
}

func NewService(repo comment_i.Repository, mapper comment_i.Mapper) Service {
	return Service{repo, mapper}
}

func (s Service) List(req comment_i.ListRequest) (comment_i.ListResponse, error) {
	var rows []comment_i.Model
	filter := s.mapper.ListRequestToFilter(req)
	err := s.repo.List(&rows, filter)
	resp := s.mapper.ModelToListResponse(rows)
	
	return resp, err
}

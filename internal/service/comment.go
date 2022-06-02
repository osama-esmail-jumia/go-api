package service

import (
	"go-api/internal/mapper"
	"go-api/internal/model"
	"go-api/internal/repository"
	"go-api/internal/request"
	"go-api/internal/response"
)

type CommentI interface {
	List(req request.CommentList) (response.CommentList, error)
}

type Comment struct {
	repo   repository.CommentI
	mapper mapper.CommentI
}

func NewComment(repo repository.CommentI, mapper mapper.CommentI) Comment {
	return Comment{repo, mapper}
}

func (s Comment) List(req request.CommentList) (response.CommentList, error) {
	var rows []model.Comment
	filter := s.mapper.ListRequestToFilter(req)
	err := s.repo.List(&rows, filter)
	resp := s.mapper.ModelToListResponse(rows)
	
	return resp, err
}

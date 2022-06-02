package service

import (
	mapper_i "go-api/definitions/mapper"
	"go-api/definitions/model"
	"go-api/definitions/repository"
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Comment struct {
	repo   repository_i.Comment
	mapper mapper_i.Comment
}

func NewComment(repo repository_i.Comment, mapper mapper_i.Comment) Comment {
	return Comment{repo, mapper}
}

func (s Comment) List(req request.CommentList) (response.CommentList, error) {
	var rows []model.Comment
	filter := s.mapper.ListRequestToFilter(req)
	err := s.repo.List(&rows, filter)
	resp := s.mapper.ModelToListResponse(rows)
	
	return resp, err
}

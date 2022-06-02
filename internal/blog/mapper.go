package blog

import (
	"go-api/definitions/blog"
	"go-api/definitions/comment"
)

type Mapper struct {
	commentMapper comment_i.Mapper
}

func NewMapper(commentMapper comment_i.Mapper) Mapper {
	return Mapper{commentMapper}
}

func (m Mapper) CreateRequestToModel(req blog_i.CreateRequest) blog_i.Model {
	return blog_i.Model{
		Title:       &req.Title,
		Description: &req.Description,
	}
}

func (m Mapper) UpdateRequestToModel(req blog_i.UpdateRequest) blog_i.Model {
	return blog_i.Model{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}
}

func (m Mapper) DeleteRequestToModel(req blog_i.DeleteRequest) blog_i.Model {
	return blog_i.Model{
		ID: req.ID,
	}
}

func (m Mapper) ModelToResponse(model blog_i.Model) blog_i.Response {
	resp := blog_i.Response{
		ID: model.ID,
	}
	
	if model.Title != nil {
		resp.Title = *model.Title
	}
	
	if model.Description != nil {
		resp.Description = *model.Description
	}
	
	resp.Comments = m.commentMapper.ModelToListResponse(model.Comments)
	
	return resp
}

func (m Mapper) ModelToListResponse(models []blog_i.Model) blog_i.ListResponse {
	rows := make([]blog_i.Response, 0, len(models))
	for _, v := range models {
		rows = append(rows, m.ModelToResponse(v))
	}
	
	return rows
}

func (m Mapper) ListRequestToFilter(req blog_i.ListRequest) blog_i.ModelFilter {
	return blog_i.ModelFilter{
		Title:       req.Title,
		Description: req.Description,
	}
}

package mapper

import (
	"go-api/internal/model"
	"go-api/internal/request"
	"go-api/internal/response"
)

type CommentI interface {
	ModelToResponse(model model.Comment) response.Comment
	ModelToListResponse(model []model.Comment) response.CommentList
	ListRequestToFilter(req request.CommentList) model.CommentFilter
}

type Comment struct {
	blogMapper BlogI
}

func NewComment(blogMapper BlogI) Comment {
	return Comment{blogMapper}
}

func (m Comment) ModelToResponse(model model.Comment) response.Comment {
	resp := response.Comment{
		ID:     model.ID,
		BlogID: model.BlogID,
	}
	
	if model.Text != nil {
		resp.Text = *model.Text
	}
	
	if model.Blog != nil {
		b := m.blogMapper.ModelToResponse(*model.Blog)
		resp.Blog = &b
	}
	
	return resp
}

func (m Comment) ModelToListResponse(models []model.Comment) response.CommentList {
	rows := make([]response.Comment, 0, len(models))
	for _, v := range models {
		rows = append(rows, m.ModelToResponse(v))
	}
	
	return rows
}

func (m Comment) ListRequestToFilter(req request.CommentList) model.CommentFilter {
	return model.CommentFilter{
		BlogID: req.BlogID,
		Text:   req.Text,
	}
}

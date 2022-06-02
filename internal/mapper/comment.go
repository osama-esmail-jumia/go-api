package mapper

import (
	"go-api/definitions/http_model"
	"go-api/definitions/mapper"
	"go-api/definitions/model"
)

type Comment struct {
	blogMapper mapper_i.Blog
}

func NewComment(blogMapper mapper_i.Blog) Comment {
	return Comment{blogMapper}
}

func (m Comment) ModelToResponse(model model.Comment) http_model.CommentResponse {
	resp := http_model.CommentResponse{
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

func (m Comment) ModelToListResponse(models []model.Comment) http_model.CommentListResponse {
	rows := make([]http_model.CommentResponse, 0, len(models))
	for _, v := range models {
		rows = append(rows, m.ModelToResponse(v))
	}
	
	return rows
}

func (m Comment) ListRequestToFilter(req http_model.CommentListRequest) model.CommentFilter {
	return model.CommentFilter{
		BlogID: req.BlogID,
		Text:   req.Text,
	}
}

package mapper

import (
	"go-api/definitions/http_model"
	"go-api/definitions/mapper"
	"go-api/definitions/model"
)

type Blog struct {
	commentMapper mapper_i.Comment
}

func NewBlog(commentMapper mapper_i.Comment) Blog {
	return Blog{commentMapper}
}

func (m Blog) CreateRequestToModel(req http_model.BlogCreateRequest) model.Blog {
	return model.Blog{
		Title:       &req.Title,
		Description: &req.Description,
	}
}

func (m Blog) UpdateRequestToModel(req http_model.BlogUpdateRequest) model.Blog {
	return model.Blog{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}
}

func (m Blog) DeleteRequestToModel(req http_model.BlogDeleteRequest) model.Blog {
	return model.Blog{
		ID: req.ID,
	}
}

func (m Blog) ModelToResponse(model model.Blog) http_model.BlogResponse {
	resp := http_model.BlogResponse{
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

func (m Blog) ModelToListResponse(models []model.Blog) http_model.BlogListResponse {
	rows := make([]http_model.BlogResponse, 0, len(models))
	for _, v := range models {
		rows = append(rows, m.ModelToResponse(v))
	}
	
	return rows
}

func (m Blog) ListRequestToFilter(req http_model.BlogListRequest) model.BlogFilter {
	return model.BlogFilter{
		Title:       req.Title,
		Description: req.Description,
	}
}

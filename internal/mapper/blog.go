package mapper

import (
	"go-api/definitions/mapper"
	"go-api/definitions/model"
	"go-api/definitions/request"
	"go-api/definitions/response"
)

type Blog struct {
	commentMapper mapper.Comment
}

func NewBlog(commentMapper mapper.Comment) Blog {
	return Blog{commentMapper}
}

func (m Blog) CreateRequestToModel(req request.BlogCreate) model.Blog {
	return model.Blog{
		Title:       &req.Title,
		Description: &req.Description,
	}
}

func (m Blog) UpdateRequestToModel(req request.BlogUpdate) model.Blog {
	return model.Blog{
		ID:          req.ID,
		Title:       req.Title,
		Description: req.Description,
	}
}

func (m Blog) DeleteRequestToModel(req request.BlogDelete) model.Blog {
	return model.Blog{
		ID: req.ID,
	}
}

func (m Blog) ModelToResponse(model model.Blog) response.Blog {
	resp := response.Blog{
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

func (m Blog) ModelToListResponse(models []model.Blog) response.BlogList {
	rows := make([]response.Blog, 0, len(models))
	for _, v := range models {
		rows = append(rows, m.ModelToResponse(v))
	}
	
	return rows
}

func (m Blog) ListRequestToFilter(req request.BlogList) model.BlogFilter {
	return model.BlogFilter{
		Title:       req.Title,
		Description: req.Description,
	}
}

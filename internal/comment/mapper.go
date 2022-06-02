package comment

import (
	"go-api/definitions/blog"
	"go-api/definitions/comment"
)

type Mapper struct {
	blogMapper blog_i.Mapper
}

func NewMapper(blogMapper blog_i.Mapper) Mapper {
	return Mapper{blogMapper}
}

func (m Mapper) ModelToResponse(model comment_i.Model) comment_i.Response {
	resp := comment_i.Response{
		ID:     model.ID,
		BlogID: model.BlogID,
	}
	
	if model.Text != nil {
		resp.Text = *model.Text
	}
	
	// Circular dependency error
	//if model.Blog != nil {
	//	b := m.blogMapper.ModelToResponse(*model.Blog)
	//	resp.Blog = &b
	//}
	
	return resp
}

func (m Mapper) ModelToListResponse(models []comment_i.Model) comment_i.ListResponse {
	rows := make([]comment_i.Response, 0, len(models))
	for _, v := range models {
		rows = append(rows, m.ModelToResponse(v))
	}
	
	return rows
}

func (m Mapper) ListRequestToFilter(req comment_i.ListRequest) comment_i.ModelFilter {
	return comment_i.ModelFilter{
		BlogID: req.BlogID,
		Text:   req.Text,
	}
}

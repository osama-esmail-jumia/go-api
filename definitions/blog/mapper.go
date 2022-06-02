package blog_i

type Mapper interface {
	CreateRequestToModel(req CreateRequest) Model
	UpdateRequestToModel(req UpdateRequest) Model
	DeleteRequestToModel(req DeleteRequest) Model
	ModelToResponse(model Model) Response
	ModelToListResponse(model []Model) ListResponse
	ListRequestToFilter(req ListRequest) ModelFilter
}

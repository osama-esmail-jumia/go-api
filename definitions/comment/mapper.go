package comment_i

type Mapper interface {
	ModelToResponse(model Model) Response
	ModelToListResponse(model []Model) ListResponse
	ListRequestToFilter(req ListRequest) ModelFilter
}

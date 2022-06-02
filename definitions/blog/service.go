package blog_i

type Service interface {
	Create(create CreateRequest) (Response, error)
	Update(req UpdateRequest) (Response, error)
	Delete(req DeleteRequest) error
	List(req ListRequest) (ListResponse, error)
}

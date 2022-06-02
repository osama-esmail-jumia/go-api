package comment_i

type Service interface {
	List(req ListRequest) (ListResponse, error)
}

package comment_i

type Repository interface {
	List(models *[]Model, filter ModelFilter) error
}

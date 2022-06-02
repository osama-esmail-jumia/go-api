package blog_i

type Repository interface {
	Create(model *Model) error
	Delete(model *Model) error
	Update(model *Model) error
	List(models *[]Model, filter ModelFilter) error
}

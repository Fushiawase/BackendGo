package document

type Repository interface {
	Create(d *Document) error

	FindByID(id int) (*Document, error)

	Delete(id int) error
}

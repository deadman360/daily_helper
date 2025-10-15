package infrastructure

type repository struct {
	basePath string
}

type IRepository interface {
	Create(string)
	// Update()
	// Delete()
	// FindAll()
	// FindByName()
	// FindByDate()
}

func NewRepository(basePath string) IRepository {
	return &repository{
		basePath: basePath,
	}
}

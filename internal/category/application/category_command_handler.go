package application

type CategoryCommandHandler struct {
	categoryRepository CategoryRepository
}

func NewCategoryCommandHandler(repository CategoryRepository) *CategoryCommandHandler {
	return &CategoryCommandHandler{categoryRepository: repository}
}

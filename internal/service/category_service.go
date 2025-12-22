package service

import (
	"inventory-system/internal/entity"
	"inventory-system/internal/repository"
)

type CategoryService struct {
	categoryRepository repository.ItemCategoryRepository
}

func NewCategoryService(categoryRepository repository.ItemCategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository}
}

func (service *CategoryService) GetAll() ([]entity.ItemCategory, error) {
	return service.categoryRepository.FindAll()
}

func (service *CategoryService) Create(name string) error {
	category := entity.ItemCategory{Name: name}
	return service.categoryRepository.Save(category)
}

func (service *CategoryService) Update(id int, name string) error {
	return service.categoryRepository.Update(entity.ItemCategory{
		ID:   id,
		Name: name,
	})
}

func (service *CategoryService) Delete(id int) error {
	return service.categoryRepository.Delete(id)
}

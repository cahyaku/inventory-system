package service

import (
	"errors"
	"inventory-system/internal/entity"
	"inventory-system/internal/repository"
	"strings"
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
	if name == "" {
		return errors.New("category name cannot be empty")
	}

	categories, err := service.categoryRepository.FindAll()
	if err != nil {
		return err
	}

	for _, cat := range categories {
		if strings.EqualFold(cat.Name, name) {
			return errors.New("category already exists")
		}
	}

	category := entity.ItemCategory{Name: name}
	return service.categoryRepository.Save(category)
}

func (service *CategoryService) Update(id int, name string) error {
	if name == "" {
		return errors.New("category name cannot be empty")
	}

	// pastikan categori ada
	currentCategory, err := service.categoryRepository.FindByID(id)
	if err != nil {
		return err
	}

	// jika nama tidak berubah (bisa)
	if strings.EqualFold(currentCategory.Name, name) {
		return nil
	}

	// periksa apakah nama sudah ada
	categories, err := service.categoryRepository.FindAll()
	if err != nil {
		return err
	}

	for _, cat := range categories {
		if strings.EqualFold(cat.Name, name) {
			return errors.New("category already exists")
		}
	}

	return service.categoryRepository.Update(entity.ItemCategory{
		ID:   id,
		Name: name,
	})
}

func (service *CategoryService) Delete(id int) error {
	return service.categoryRepository.Delete(id)
}

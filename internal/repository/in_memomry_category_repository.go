package repository

import (
	"errors"
	"inventory-system/internal/entity"
)

type InMemoryCategoryRepository struct {
	categories []entity.ItemCategory
	nextID     int
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	// dummy data agar create item bisa jalan
	return &InMemoryCategoryRepository{
		categories: []entity.ItemCategory{
			//{ID: 1, Name: "Electronics"},
			//{ID: 2, Name: "Books"},
		},
	}
}

func (repo *InMemoryCategoryRepository) FindAll() ([]entity.ItemCategory, error) {
	return repo.categories, nil
}

func (repo *InMemoryCategoryRepository) FindByID(id int) (entity.ItemCategory, error) {
	for _, category := range repo.categories {
		if category.ID == id {
			return category, nil
		}
	}
	return entity.ItemCategory{}, errors.New("Category not found!")
}

func (repo *InMemoryCategoryRepository) Save(category entity.ItemCategory) error {
	category.ID = repo.NextID()
	repo.nextID++
	repo.categories = append(repo.categories, category)
	return nil
}

func (repo *InMemoryCategoryRepository) Update(category entity.ItemCategory) error {
	for i, cat := range repo.categories {
		if cat.ID == category.ID {
			repo.categories[i] = category
			return nil
		}
	}
	return errors.New("Category not found!")
}

func (repo *InMemoryCategoryRepository) Delete(id int) error {
	for i, cat := range repo.categories {
		if cat.ID == id {
			repo.categories = append(repo.categories[:i], repo.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("Category not found!")
}

func (repo *InMemoryCategoryRepository) NextID() int {
	repo.nextID++
	return repo.nextID
}

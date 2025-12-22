package repository

import "inventory-system/internal/entity"

type InMemoryCategoryRepository struct {
	categories []entity.ItemCategory
	nextID     int
}

func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	// dummy data agar create item bisa jalan
	return &InMemoryCategoryRepository{
		categories: []entity.ItemCategory{
			{ID: 1, Name: "Electronics"},
			{ID: 2, Name: "Books"},
		},
	}
}

func (repo *InMemoryCategoryRepository) FindAll() ([]entity.ItemCategory, error) {
	return repo.categories, nil
}

func (repo *InMemoryCategoryRepository) FindByID(id int) (entity.ItemCategory, error) {
	return entity.ItemCategory{}, nil
}

func (repo *InMemoryCategoryRepository) Save(category entity.ItemCategory) error {
	repo.categories = append(repo.categories, category)
	return nil
}

func (repo *InMemoryCategoryRepository) Update(category entity.ItemCategory) error {
	return nil
}

func (repo *InMemoryCategoryRepository) Delete(id int) error {
	return nil
}

func (repo *InMemoryCategoryRepository) NextID() int {
	repo.nextID++
	return repo.nextID
}

func (repo *InMemoryCategoryRepository) ResetID() {
	repo.nextID = 0
}

func (repo *InMemoryCategoryRepository) FindByName(name string) (entity.ItemCategory, error) {
	return entity.ItemCategory{}, nil
}

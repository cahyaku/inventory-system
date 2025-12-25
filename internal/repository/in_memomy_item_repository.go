package repository

import (
	"errors"
	"inventory-system/internal/entity"
)

type InMemoryItemRepository struct {
	items  []entity.Item
	nextID int
}

func NewInMemoryItemRepository() *InMemoryItemRepository {
	return &InMemoryItemRepository{
		items:  []entity.Item{},
		nextID: 1,
	}
}

func (repo *InMemoryItemRepository) FindAll() ([]entity.Item, error) {
	return repo.items, nil
}

func (repo *InMemoryItemRepository) FindByID(id int) (entity.Item, error) {
	for _, item := range repo.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Item{}, errors.New("item not found")
}

func (repo *InMemoryItemRepository) Save(item entity.Item) error {
	item.ID = repo.nextID
	repo.nextID++
	repo.items = append(repo.items, item)
	return nil
}

func (repo *InMemoryItemRepository) Update(item entity.Item) error {
	//repo.items[item.ID-1] = item
	for i, it := range repo.items {
		if it.ID == item.ID {
			repo.items[i] = item
			return nil
		}
		return errors.New("item not found")
	}
	return nil
}

func (repo *InMemoryItemRepository) Delete(id int) error {
	for i, item := range repo.items {
		if item.ID == id {
			repo.items = append(repo.items[:i], repo.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

func (repo *InMemoryItemRepository) CountByCategoryID(categoryID int) (int, error) {
	count := 0

	for _, item := range repo.items {
		for _, cat := range item.Categories {
			if cat.ID == categoryID {
				count++
				break // hindari double count
			}
		}
	}

	return count, nil
}

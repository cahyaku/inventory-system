package service

import (
	"errors"
	"inventory-system/internal/entity"
	"inventory-system/internal/repository"
)

type ItemService struct {
	itemRepo     repository.ItemRepository
	categoryRepo repository.ItemCategoryRepository
}

func NewItemService(itemRepo repository.ItemRepository, categoryRepo repository.ItemCategoryRepository) *ItemService {
	return &ItemService{itemRepo, categoryRepo}
}

func (service *ItemService) GetAllItems() ([]entity.Item, error) {
	return service.itemRepo.FindAll()
}

func (service *ItemService) CreateItem(
	name string,
	stock int,
	categoryIDs []int,
) error {

	// 1. validasi basic
	if name == "" {
		return errors.New("item name cannot be empty")
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	if len(categoryIDs) == 0 {
		return errors.New("item must have at least one category")
	}

	// 2. Ambil semua category
	categories, _ := service.categoryRepo.FindAll()
	if len(categories) == 0 {
		return errors.New("no categories available")
	}

	// 3. Cocokan ID kategori
	var selected []entity.ItemCategory
	for _, id := range categoryIDs {
		for _, c := range categories {
			if c.ID == id {
				selected = append(selected, c)
			}
		}
	}

	if len(selected) == 0 {
		return errors.New("item must have at least one category")
	}

	item := entity.Item{
		ID:         len(categories) + 1,
		Name:       name,
		Stock:      stock,
		Categories: selected,
	}

	return service.itemRepo.Save(item)
}

func (service *ItemService) GetByID(id int) (entity.Item, error) {
	return service.itemRepo.FindByID(id)
}

func (service *ItemService) Save(item entity.Item) error {
	return service.itemRepo.Save(item)
}

func (service *ItemService) Update(item entity.Item) error {
	return service.itemRepo.Update(item)
}

func (service *ItemService) Delete(id int) error {
	return service.itemRepo.Delete(id)
}

func (service *ItemService) DeleteAll() error {
	return nil
}

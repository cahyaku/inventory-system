package service

import (
	"errors"
	"inventory-system/internal/entity"
	"inventory-system/internal/repository"
	"strconv"
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

func (service *ItemService) Create(
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

	categoryMap := make(map[int]entity.ItemCategory)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}

	// Validasi category
	var selected []entity.ItemCategory
	seen := make(map[int]bool)

	for _, id := range categoryIDs {
		category, exists := categoryMap[id]
		if !exists {
			return errors.New("Invalid category ID: " + strconv.Itoa(id))
		}
		if seen[id] {
			return errors.New("duplicate category ID: " + strconv.Itoa(id))
		}
		seen[id] = true
		selected = append(selected, category)
	}

	if len(selected) == 0 {
		return errors.New("item must have at least one category")
	}

	item := entity.Item{
		//ID:         len(categories) + 1,
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

func (service *ItemService) Update(
	itemID int,
	name string,
	stock int,
	categoryIDs []int,
) error {

	// 1. Ambil item lama
	item, err := service.itemRepo.FindByID(itemID)
	if err != nil {
		return err
	}

	// 2. Validasi
	if name == "" {
		return errors.New("item name cannot be empty")
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	if len(categoryIDs) == 0 {
		return errors.New("item must have at least one category")
	}

	// 3. Ambil semua kategori
	categories, _ := service.categoryRepo.FindAll()
	if len(categories) == 0 {
		return errors.New("no categories available")
	}

	// 4. Map kategori untuk validasi cepat
	categoryMap := make(map[int]entity.ItemCategory)
	for _, c := range categories {
		categoryMap[c.ID] = c
	}

	var selected []entity.ItemCategory
	seen := make(map[int]bool)

	for _, id := range categoryIDs {
		cat, exists := categoryMap[id]
		if !exists {
			return errors.New("invalid category ID: " + strconv.Itoa(id))
		}
		if seen[id] {
			return errors.New("duplicate category ID: " + strconv.Itoa(id))
		}
		seen[id] = true
		selected = append(selected, cat)
	}

	// 5. Update data item
	item.Name = name
	item.Stock = stock
	item.Categories = selected

	// 6. Simpan
	return service.itemRepo.Update(item)
}

func (service *ItemService) Delete(id int) error {
	_, err := service.itemRepo.FindByID(id)
	if err != nil {
		return err
	}
	return service.itemRepo.Delete(id)
}

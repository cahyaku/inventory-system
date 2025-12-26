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

func NewItemService(
	itemRepo repository.ItemRepository,
	categoryRepo repository.ItemCategoryRepository,
) *ItemService {
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

	// Simple validation
	if name == "" {
		return errors.New("item name cannot be empty")
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	items, err := service.itemRepo.FindAll()
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.Name == name {
			return errors.New("item already exists")
		}
	}

	// reusable category validation
	categories, err := service.validateAndGetCategories(categoryIDs)
	if err != nil {
		return err
	}

	item := entity.Item{
		Name:       name,
		Stock:      stock,
		Categories: categories,
	}

	return service.itemRepo.Save(item)
}

func (service *ItemService) Update(
	itemID int,
	name string,
	stock int,
	categoryIDs []int,
) error {

	// Take data item by ID
	item, err := service.itemRepo.FindByID(itemID)
	if err != nil {
		return err
	}

	// simple validation
	if name == "" {
		return errors.New("item name cannot be empty")
	}

	if stock < 0 {
		return errors.New("stock cannot be negative")
	}

	items, err := service.itemRepo.FindAll()
	if err != nil {
		return err
	}

	for _, item := range items {
		if item.Name == name && item.ID != itemID {
			return errors.New("item already exists")
		}
	}

	// reusable category validation
	categories, err := service.validateAndGetCategories(categoryIDs)
	if err != nil {
		return err
	}

	// Update data item
	item.Name = name
	item.Stock = stock
	item.Categories = categories

	// Save data item
	return service.itemRepo.Update(item)
}

func (service *ItemService) Delete(id int) error {
	_, err := service.itemRepo.FindByID(id)
	if err != nil {
		return err
	}
	return service.itemRepo.Delete(id)
}

// validateAndGetCategories mengambil kategori berdasarkan ID dan memvalidasi
// validateAndGetCategories (categoryIDs []int) memvalidasi ID kategori yang dipilih user.
// []entity.ItemCategory adalah outputnya.
func (service *ItemService) validateAndGetCategories(categoryIDs []int) ([]entity.ItemCategory, error) {
	if len(categoryIDs) == 0 {
		return nil, errors.New("item must have at least one category")
	}

	categories, err := service.categoryRepo.FindAll()
	if err != nil {
		return nil, err
	}

	// validasi kategori ada di sistem
	if len(categories) == 0 {
		return nil, errors.New("no categories available")
	}

	// Map untuk menyimpan kategori berdasarkan ID
	categoryMap := make(map[int]entity.ItemCategory)
	for _, cat := range categories {
		categoryMap[cat.ID] = cat
	}

	// Menyiapkan variabel hasil dan penanda duplikat
	var selected []entity.ItemCategory // menyimpan kategory yang valid
	seen := make(map[int]bool)         // menandai id yang usdah diproses

	// loop dan validasi setiap ID dari user
	for _, id := range categoryIDs {
		category, exists := categoryMap[id]
		// jika user memilih ID kategori yang tidak ada
		if !exists {
			return nil, errors.New("Invalid category ID: " + strconv.Itoa(id))
		}
		// jika user memilih ID kategori yang sudah ada
		if seen[id] {
			return nil, errors.New("duplicate category ID: " + strconv.Itoa(id))
		}
		seen[id] = true                       // tandai ID sudah dipakai
		selected = append(selected, category) // simpan kategori
	}
	return selected, nil
}

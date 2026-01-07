package service

import (
	"errors"
	"inventory-system/internal/entity"
	"inventory-system/internal/repository"
	"strings"
)

type CategoryService struct {
	categoryRepository repository.ItemCategoryRepository
	itemRepository     repository.ItemRepository // ini dipakai
	// kalau sudah dipakai, maka tidak bisa diubah atau dihapus.
}

// ini nambah itemRepository untuk cek apakah categori udah dipakai pada items?
func NewCategoryService(
	categoryRepository repository.ItemCategoryRepository,
	itemRepository repository.ItemRepository,
) *CategoryService {
	return &CategoryService{categoryRepository, itemRepository}
}

// GetAll mengambil semua kategori
func (service *CategoryService) GetAll() ([]entity.ItemCategory, error) {
	return service.categoryRepository.FindAll()
}

// Create menambahkan kategori baru
func (service *CategoryService) Create(name string) error {
	if name == "" {
		return errors.New("category name cannot be empty")
	}

	categories, err := service.categoryRepository.FindAll()
	if err != nil {
		return err
	}

	for _, cat := range categories {
		// strings.EqualFold() membuat perbandingan tidak case-insensitive
		// ex. "Food" == "food" == "FOOD"
		if strings.EqualFold(cat.Name, name) {
			return errors.New("category already exists")
		}
	}

	category := entity.ItemCategory{Name: name}
	return service.categoryRepository.Save(category)
}

// Update mengubah nama kategori
func (service *CategoryService) Update(id int, name string) error {
	if name == "" {
		return errors.New("category name cannot be empty")
	}

	// cek apakah kategori masih bisa diubah
	if !service.CanModify(id) {
		return errors.New("category cannot be edited because it is used by items")
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

	// (agar kategori tidak punya nama sama)
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

// Delete menghapus kategori
func (service *CategoryService) Delete(id int) error {
	// cek apakah kategori masih bisa dihapus
	if !service.CanModify(id) {
		return errors.New("category cannot be edited because it is used by items")
	}

	return service.categoryRepository.Delete(id)
}

// CanModify meminta hasil hitung ke repository, mengambil angka count
// kemudian membandingkan hasilnya.
// Initinya untuk cek apakah categori sudah dipakai pada item
func (service *CategoryService) CanModify(categoryID int) bool {
	count, _ := service.itemRepository.CountByCategoryID(categoryID)
	return count == 0 // Jika countnya 0 = true (category bisa di edit / delete)
}

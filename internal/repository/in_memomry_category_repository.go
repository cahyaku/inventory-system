package repository

import (
	"errors"
	"inventory-system/internal/entity"
)

// InMemoryCategoryRepository adalah repository untuk kategori item yang disimpan di memory
type InMemoryCategoryRepository struct {
	categories []entity.ItemCategory // tempat data disimpan
	nextID     int
}

// NewInMemoryCategoryRepository inisialisasi repository kategori barang (constructor)
func NewInMemoryCategoryRepository() *InMemoryCategoryRepository {
	// dummy data agar create item bisa jalan
	return &InMemoryCategoryRepository{
		categories: []entity.ItemCategory{
			//{ID: 1, Name: "Electronics"},
			//{ID: 2, Name: "Books"},
		},
	}
}

// FindAll mengambil seluruh isi repo.categories dan mengembalikan slice tersebut.
func (repo *InMemoryCategoryRepository) FindAll() ([]entity.ItemCategory, error) {
	return repo.categories, nil
}

// FindByID mengambil kategori berdasarkan ID
// Loop berdasarkan ID dan jika ID sama, maka kembalikan kategori tersebut.
func (repo *InMemoryCategoryRepository) FindByID(id int) (entity.ItemCategory, error) {
	for _, category := range repo.categories {
		if category.ID == id {
			return category, nil
		}
	}
	return entity.ItemCategory{}, errors.New("Category not found!")
}

// Save menyimpan kategori baru ke repo.categories
func (repo *InMemoryCategoryRepository) Save(category entity.ItemCategory) error {
	category.ID = repo.NextID() // id + 1
	repo.categories = append(repo.categories, category)
	return nil
}

// Update mengupdate kategori yang sudah ada
func (repo *InMemoryCategoryRepository) Update(category entity.ItemCategory) error {
	for i, cat := range repo.categories {
		if cat.ID == category.ID {
			repo.categories[i] = category // Ganti elemen slice di index i
			return nil
		}
	}
	return errors.New("Category not found!")
}

// Delete menghapus kategori berdasarkan ID
func (repo *InMemoryCategoryRepository) Delete(id int) error {
	for i, cat := range repo.categories { // loop
		if cat.ID == id { // jika id sama
			// hapus elemen slice di index i
			// repo.categories[:i] => sebelum data di index i
			// repo.categories[i+1:] => setelah data di index i
			repo.categories = append(repo.categories[:i], repo.categories[i+1:]...)
			return nil
		}
	}
	return errors.New("Category not found!")
}

// NextID mengambil ID terakhir dan menambahkan 1
func (repo *InMemoryCategoryRepository) NextID() int {
	repo.nextID++
	return repo.nextID
}

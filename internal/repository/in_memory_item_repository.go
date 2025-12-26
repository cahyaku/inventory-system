package repository

import (
	"errors"
	"inventory-system/internal/entity"
)

// InMemoryItemRepository adalah repository untuk item yang disimpan di memory
type InMemoryItemRepository struct {
	items  []entity.Item
	nextID int
}

// NewInMemoryItemRepository inisialisasi repository item (constructor)
func NewInMemoryItemRepository() ItemRepository {
	return &InMemoryItemRepository{
		items:  []entity.Item{},
		nextID: 1,
	}
}

// FindAll mengambil seluruh isi repo.items dan mengembalikan slice tersebut.
func (repo *InMemoryItemRepository) FindAll() ([]entity.Item, error) {
	return repo.items, nil
}

// FindByID mengambil item berdasarkan ID
func (repo *InMemoryItemRepository) FindByID(id int) (entity.Item, error) {
	for _, item := range repo.items {
		if item.ID == id {
			return item, nil
		}
	}
	return entity.Item{}, errors.New("item not found")
}

// Save menyimpan item baru ke repo.items
func (repo *InMemoryItemRepository) Save(item entity.Item) error {
	item.ID = repo.nextID
	repo.nextID++
	repo.items = append(repo.items, item)
	return nil
}

// Update mengupdate item yang sudah ada, dengan ID yang sama
func (repo *InMemoryItemRepository) Update(item entity.Item) error {
	for i, it := range repo.items {
		if it.ID == item.ID {
			repo.items[i] = item
			return nil
		}
	}
	return errors.New("item not found")
}

// Delete menghapus item berdasarkan ID
func (repo *InMemoryItemRepository) Delete(id int) error {
	for i, item := range repo.items {
		if item.ID == id {
			// hapus elemen slice di index i
			repo.items = append(repo.items[:i], repo.items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}

// CountByCategoryID menghitung jumlah item berdasarkan ID kategori
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

package repository

import "inventory-system/internal/entity"

// ItemRepository adalah interface untuk repository untuk item
type ItemRepository interface {
	FindAll() ([]entity.Item, error)
	FindByID(id int) (entity.Item, error)
	Save(item entity.Item) error
	Update(item entity.Item) error
	Delete(id int) error

	// Ini untuk menghitung berapa item yang menggunakan kategori
	// Perlu karena kalau kategori sudah di pakai semestinya dia tidak boleh
	// diubah atau di delete.
	CountByCategoryID(categoryID int) (int, error)
}

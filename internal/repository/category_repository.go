package repository

import "inventory-system/internal/entity"

type ItemCategoryRepository interface {
	FindAll() ([]entity.ItemCategory, error)
	FindByID(id int) (entity.ItemCategory, error)
	Save(category entity.ItemCategory) error
	Update(category entity.ItemCategory) error
	Delete(id int) error
}

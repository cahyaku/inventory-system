package repository

import "inventory-system/internal/entity"

type ItemRepository interface {
	FindAll() ([]entity.Item, error)
	FindByID(id int) (entity.Item, error)
	Save(item entity.Item) error
	Update(item entity.Item) error
	Delete(id int) error
}

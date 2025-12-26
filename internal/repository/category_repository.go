package repository

import "inventory-system/internal/entity"

// ItemCategoryRepository adalah interface untuk repository untuk kategori item
// Jadi ini mendefiniskan apa yang bisa dilakukan ke data ItemCategory, tanpa peduli cara menyimpanya.
// ItemCategoryRepository digunakan agar service tidak bergantung pada implementasi penyimpanan tertentu,
// (Hal ini mempermudah pindah storage dari in-memory ke database) misalnya.
// Intinya ini tidak menyimpan data, tidak punya field jadi hanya mendefinisikan saja.
type ItemCategoryRepository interface {
	FindAll() ([]entity.ItemCategory, error)
	FindByID(id int) (entity.ItemCategory, error)
	Save(category entity.ItemCategory) error
	Update(category entity.ItemCategory) error
	Delete(id int) error
}

// Sehingga jika ingin implementasi database, hanya perlu mengubah implementasi ItemCategoryRepository saja.
// ex: Ini pada file item_repository_postgres.go misalnya.
//type PostgresItemCategoryRepository struct {
//	db *sql.DB
//}
//
//func (r *PostgresItemCategoryRepository) FindAll() ([]entity.ItemCategory, error) {
//	rows, err := r.db.Query("SELECT id, name FROM categories")
//	if err != nil {
//		return nil, err
//	}
//
//	var categories []entity.ItemCategory
//	for rows.Next() {
//		var c entity.ItemCategory
//		rows.Scan(&c.ID, &c.Name)
//		categories = append(categories, c)
//	}
//	return categories, nil
//}

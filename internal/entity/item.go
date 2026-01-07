package entity

// Item Entity untuk merepresentasikan barang/item
// Digunakan pada layer repository, service dan controller.
type Item struct {
	ID         int
	Name       string
	Stock      int
	Categories []ItemCategory // Slice category
}

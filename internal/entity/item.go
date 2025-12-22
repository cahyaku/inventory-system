package entity

type Item struct {
	ID         int
	Name       string
	Stock      int
	Categories []ItemCategory // Slice category
}

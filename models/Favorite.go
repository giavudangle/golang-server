package models

type Favorite struct {
	ID     uint
	UserID uint
	items  []Product
}

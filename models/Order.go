package models

type Order struct {
	ID     uint
	UserID uint
	items  []Product
}

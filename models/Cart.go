package models

type Cart struct {
	ID       uint
	items    []Product
	quantity uint
	UserID   uint
}

package models

type Book struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"titlte" gorm:""`
	Author string `json:"author"`
}

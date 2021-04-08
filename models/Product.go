package models

type Product struct {
	ID            uint
	title         string
	color         string
	price         float64
	origin        string
	destination   string
	standard      string
	description   string
	imageUrl      string
	thumbImageUrl string
	productType   string
}

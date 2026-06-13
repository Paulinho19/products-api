package model

type Product struct {
	ID    int     `json:"product_id"` //specify json key name
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
package model

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Decription  string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	Category_ID int     `json:"category_id"`
}

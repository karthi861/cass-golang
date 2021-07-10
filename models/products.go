package models

type Product struct {
	ID          int     `json:"ID"`
	Name        string  `json:"Name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
}

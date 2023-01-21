package models

type Product struct {
	Id     string  `json:"id,omitempty"`
	Title  string  `json:"title,omitempty"`
	Price  float32 `json:"price,omitempty"`
	Info   string  `json:"info,omitempty"`
	Rating float32 `json:"rating,omitempty"`
}

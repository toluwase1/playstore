package models

type Product struct {
	Image       string  `json:"img"`
	ImagAlt     string  `json:"imgalt"`
	Price       float64 `json:"price"`
	Promotion   float64 `json:"promotion"`
	ProductName string  `json:"productname"`
	Description string  `json:"desc"`
}

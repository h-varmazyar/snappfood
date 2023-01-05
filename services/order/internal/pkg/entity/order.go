package entity

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	OrderID int64  `json:"order_id"`
	Price   int64  `json:"price"`
	Title   string `json:"title"`
}

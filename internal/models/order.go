package models

import (
	"time"
)

type Order struct {
	ID          string      `bson:"_id" json:"id"`
	User_ID     string      `bson:"user_id" json:"user_id"`
	Order_Items []OrderItem `bson:"order_items" json:"order_items"`
	Total       float64     `bson:"total" json:"total"`
	Status      string      `bson:"status" json:"status"` // "pending", "shipped", "delivered"
	Created_At  time.Time   `bson:"created_at" json:"created_at"`
	Updated_At  time.Time   `bson:"updated_at" json:"updated_at"`
}

type OrderItem struct {
	Product_ID string  `bson:"product_id" json:"product_id"`
	Quantity   int     `bson:"quantity" json:"quantity"`
	Price      float64 `bson:"price" json:"price"`
}

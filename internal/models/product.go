package models

import (
	"time"
)

type Product struct {
	ID          string    `bson:"_id" json:"id"`
	Name        string    `bson:"name" json:"name"`
	Description string    `bson:"description" json:"description"`
	Price       float64   `bson:"price" json:"price"`
	Category    string    `bson:"category" json:"category"`
	Stock       int       `bson:"stock" json:"stock"`
	Created_At  time.Time `bson:"created_at" json:"created_at"`
	Updated_At  time.Time `bson:"updated_at" json:"updated_at"`
}

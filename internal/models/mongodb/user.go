package mongodb

import (
	"time"
)

type User struct {
	ID         string    `bson:"_id" json:"id"`
	Name       string    `bson:"name" json:"name"`
	Email      string    `bson:"email" json:"email"`
	Created_At time.Time `bson:"created_at" json:"created_at"`
	Updated_At time.Time `bson:"updated_at" json:"updated_at"`

	Profile   Profile   `bson:"profile" json:"profile"`
	Addresses []Address `bson:"addresses" json:"addresses"`
}

type Profile struct {
	Phone         string                 `bson:"phone" json:"phone"`
	Date_Of_Birth *time.Time             `bson:"date_of_birth" json:"date_of_birth"`
	Preferences   map[string]interface{} `bson:"preferences" json:"preferences"`
}

type Address struct {
	ID      string `bson:"id" json:"id"`
	Type    string `bson:"type" json:"type"` // "billing", "shipping"
	Street  string `bson:"street" json:"street"`
	City    string `bson:"city" json:"city"`
	State   string `bson:"state" json:"state"`
	Zip     string `bson:"zip" json:"zip"`
	Country string `bson:"country" json:"country"`
}

package main

import (
	"api-servers/internal/models/mongodb"
	mongodb_repo "api-servers/internal/repository/mongodb"
	"context"
	"log"
	"time"

	"github.com/google/uuid"
)

func seed_mongodb() error {
	db, err := mongodb_repo.Connect()
	if err != nil {
		return err
	}
	defer db.Disconnect()

	ctx := context.Background()

	users := create_sample_users()
	err = seed_users(db, ctx, users)
	if err != nil {
		return err
	}

	products := create_sample_products()
	err = seed_products(db, ctx, products)
	if err != nil {
		return err
	}

	orders := create_sample_orders(users, products)
	err = seed_orders(db, ctx, orders)
	if err != nil {
		return err
	}

	log.Printf("created %d users, %d products, %d orders", len(users), len(products), len(orders))
	return nil
}

func create_sample_users() []mongodb.User {
	now := time.Now()
	birth_date_1985 := time.Date(1985, 6, 15, 0, 0, 0, 0, time.UTC)
	birth_date_1990 := time.Date(1990, 3, 22, 0, 0, 0, 0, time.UTC)

	return []mongodb.User{
		{
			ID:         uuid.New().String(),
			Name:       "Alice Johnson",
			Email:      "alice.johnson@email.com",
			Created_At: now,
			Updated_At: now,
			Profile: mongodb.Profile{
				Phone:         "+1-555-0123",
				Date_Of_Birth: &birth_date_1985,
				Preferences: map[string]interface{}{
					"newsletter":    true,
					"theme":         "dark",
					"notifications": true,
					"language":      "en",
				},
			},
			Addresses: []mongodb.Address{
				{
					ID:      uuid.New().String(),
					Type:    "home",
					Street:  "123 Oak Street",
					City:    "Portland",
					State:   "OR",
					Zip:     "97201",
					Country: "USA",
				},
				{
					ID:      uuid.New().String(),
					Type:    "billing",
					Street:  "456 Work Plaza",
					City:    "Portland",
					State:   "OR",
					Zip:     "97205",
					Country: "USA",
				},
			},
		},
		{
			ID:         uuid.New().String(),
			Name:       "Bob Smith",
			Email:      "bob.smith@email.com",
			Created_At: now,
			Updated_At: now,
			Profile: mongodb.Profile{
				Phone:         "+1-555-0456",
				Date_Of_Birth: &birth_date_1990,
				Preferences: map[string]interface{}{
					"newsletter":    false,
					"theme":         "light",
					"notifications": false,
					"language":      "en",
				},
			},
			Addresses: []mongodb.Address{
				{
					ID:      uuid.New().String(),
					Type:    "home",
					Street:  "789 Pine Avenue",
					City:    "Seattle",
					State:   "WA",
					Zip:     "98101",
					Country: "USA",
				},
			},
		},
		{
			ID:         uuid.New().String(),
			Name:       "Carol Davis",
			Email:      "carol.davis@email.com",
			Created_At: now,
			Updated_At: now,
			Profile: mongodb.Profile{
				Phone:         "+1-555-0789",
				Date_Of_Birth: nil,
				Preferences: map[string]interface{}{
					"newsletter":    true,
					"theme":         "auto",
					"notifications": true,
					"language":      "es",
				},
			},
			Addresses: []mongodb.Address{
				{
					ID:      uuid.New().String(),
					Type:    "home",
					Street:  "321 Elm Drive",
					City:    "San Francisco",
					State:   "CA",
					Zip:     "94102",
					Country: "USA",
				},
			},
		},
	}
}

func create_sample_products() []mongodb.Product {
	now := time.Now()

	return []mongodb.Product{
		{
			ID:          uuid.New().String(),
			Name:        "Wireless Bluetooth Headphones",
			Description: "High-quality wireless headphones with noise cancellation and 30-hour battery life.",
			Price:       99.99,
			Category:    "Electronics",
			Stock:       50,
			Created_At:  now,
			Updated_At:  now,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Organic Coffee Beans - Dark Roast",
			Description: "Premium organic coffee beans, dark roast, sourced from sustainable farms.",
			Price:       24.99,
			Category:    "Food & Beverage",
			Stock:       200,
			Created_At:  now,
			Updated_At:  now,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Ergonomic Office Chair",
			Description: "Comfortable ergonomic office chair with lumbar support and adjustable height.",
			Price:       299.99,
			Category:    "Furniture",
			Stock:       25,
			Created_At:  now,
			Updated_At:  now,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Stainless Steel Water Bottle",
			Description: "Insulated stainless steel water bottle that keeps drinks cold for 24 hours.",
			Price:       19.99,
			Category:    "Sports & Outdoors",
			Stock:       100,
			Created_At:  now,
			Updated_At:  now,
		},
		{
			ID:          uuid.New().String(),
			Name:        "Programming Books Bundle",
			Description: "Collection of 5 programming books covering Go, JavaScript, Python, and system design.",
			Price:       149.99,
			Category:    "Books",
			Stock:       75,
			Created_At:  now,
			Updated_At:  now,
		},
	}
}

func create_sample_orders(users []mongodb.User, products []mongodb.Product) []mongodb.Order {
	now := time.Now()

	return []mongodb.Order{
		{
			ID:      uuid.New().String(),
			User_ID: users[0].ID,
			Order_Items: []mongodb.OrderItem{
				{
					Product_ID: products[0].ID,
					Quantity:   1,
					Price:      99.99,
				},
				{
					Product_ID: products[1].ID,
					Quantity:   2,
					Price:      24.99,
				},
			},
			Total:      149.97,
			Status:     "delivered",
			Created_At: now.AddDate(0, 0, -5),
			Updated_At: now.AddDate(0, 0, -2),
		},
		{
			ID:      uuid.New().String(),
			User_ID: users[1].ID,
			Order_Items: []mongodb.OrderItem{
				{
					Product_ID: products[2].ID,
					Quantity:   1,
					Price:      299.99,
				},
			},
			Total:      299.99,
			Status:     "shipped",
			Created_At: now.AddDate(0, 0, -3),
			Updated_At: now.AddDate(0, 0, -1),
		},
		{
			ID:      uuid.New().String(),
			User_ID: users[2].ID,
			Order_Items: []mongodb.OrderItem{
				{
					Product_ID: products[3].ID,
					Quantity:   3,
					Price:      19.99,
				},
				{
					Product_ID: products[4].ID,
					Quantity:   1,
					Price:      149.99,
				},
			},
			Total:      209.96,
			Status:     "pending",
			Created_At: now.AddDate(0, 0, -1),
			Updated_At: now,
		},
	}
}

func seed_users(db *mongodb_repo.Database, ctx context.Context, users []mongodb.User) error {
	collection := db.Database.Collection("users")

	for _, user := range users {
		_, err := collection.InsertOne(ctx, user)
		if err != nil {
			return err
		}
		log.Printf("created user: %s (%s)", user.Name, user.Email)
	}
	return nil
}

func seed_products(db *mongodb_repo.Database, ctx context.Context, products []mongodb.Product) error {
	collection := db.Database.Collection("products")

	for _, product := range products {
		_, err := collection.InsertOne(ctx, product)
		if err != nil {
			return err
		}
		log.Printf("created product: %s ($%.2f)", product.Name, product.Price)
	}
	return nil
}

func seed_orders(db *mongodb_repo.Database, ctx context.Context, orders []mongodb.Order) error {
	collection := db.Database.Collection("orders")

	for _, order := range orders {
		_, err := collection.InsertOne(ctx, order)
		if err != nil {
			return err
		}
		log.Printf("created order: %s ($%.2f, %s)", order.ID[:8], order.Total, order.Status)
	}
	return nil
}
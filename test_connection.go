package main

import (
	"api-servers/internal/repository/mongodb"
	"log"
)

func main() {
	log.Println("Testing MongoDB connection...")

	// Test connection
	db, err := mongodb.Connect()
	if err != nil {
		log.Fatalf("Connection failed: %v", err)
	}

	log.Println("✅ Connection successful!")
	log.Printf("Database name: %s", db.Database.Name())

	// Test disconnect
	err = db.Disconnect()
	if err != nil {
		log.Printf("Disconnect failed: %v", err)
	} else {
		log.Println("✅ Disconnected successfully!")
	}
}

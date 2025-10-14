package mongodb

import (
	"api-servers/internal/models/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type userRepository struct {
	db *Database
}

func NewUserRepository(db *Database) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user mongodb.User) error {
	collection := r.db.Database.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	return err
}

func (r *userRepository) GetByID(id string) (mongodb.User, error) {
	collection := r.db.Database.Collection("users")
	var user mongodb.User
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

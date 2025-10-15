package mongodb

import (
	"api-servers/internal/models/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	db         *Database
	collection *mongo.Collection
}

func NewUserRepository(db *Database) UserRepository {
	return &userRepository{
		db:         db,
		collection: db.Database.Collection("users"),
	}
}

func (r *userRepository) Create(user mongodb.User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *userRepository) GetByID(id string) (mongodb.User, error) {
	var user mongodb.User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return user, err
}

func (r *userRepository) GetByEmail(email string) (mongodb.User, error) {
	var user mongodb.User
	err := r.collection.FindOne(context.Background(), bson.M{"email": email}).Decode(&user)
	return user, err
}

func (r *userRepository) GetAll() ([]mongodb.User, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var users []mongodb.User
	err = cursor.All(context.Background(), &users)
	return users, err
}

func (r *userRepository) Update(id string, user mongodb.User) error {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	return err
}

func (r *userRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

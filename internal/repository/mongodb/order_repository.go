package mongodb

import (
	"api-servers/internal/models/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type orderRepository struct {
	db         *Database
	collection *mongo.Collection
}

func NewOrderRepository(db *Database) OrderRepository {
	return &orderRepository{
		db:         db,
		collection: db.Database.Collection("orders"),
	}
}

func (r *orderRepository) Create(order mongodb.Order) error {
	_, err := r.collection.InsertOne(context.Background(), order)
	return err
}

func (r *orderRepository) GetByID(id string) (mongodb.Order, error) {
	var order mongodb.Order
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&order)
	return order, err
}

func (r *orderRepository) GetByUserID(user_id string) ([]mongodb.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"user_id": user_id})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []mongodb.Order
	err = cursor.All(context.Background(), &orders)
	return orders, err
}

func (r *orderRepository) GetByStatus(status string) ([]mongodb.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"status": status})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []mongodb.Order
	err = cursor.All(context.Background(), &orders)
	return orders, err
}

func (r *orderRepository) GetAll() ([]mongodb.Order, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var orders []mongodb.Order
	err = cursor.All(context.Background(), &orders)
	return orders, err
}

func (r *orderRepository) Update(id string, user mongodb.Order) error {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, user)
	return err
}

func (r *orderRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

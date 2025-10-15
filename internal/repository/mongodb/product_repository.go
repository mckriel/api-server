package mongodb

import (
	"api-servers/internal/models/mongodb"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type productRepository struct {
	db         *Database
	collection *mongo.Collection
}

func NewProductRepository(db *Database) ProductRepository {
	return &productRepository{
		db:         db,
		collection: db.Database.Collection("products"),
	}
}

func (r *productRepository) Create(product mongodb.Product) error {
	_, err := r.collection.InsertOne(context.Background(), product)
	return err
}

func (r *productRepository) GetByID(id string) (mongodb.Product, error) {
	var product mongodb.Product
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&product)
	return product, err
}

func (r *productRepository) GetByCategory(category string) ([]mongodb.Product, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{"category": category})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []mongodb.Product
	err = cursor.All(context.Background(), &products)
	return products, err
}

func (r *productRepository) GetAll() ([]mongodb.Product, error) {
	cursor, err := r.collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	var products []mongodb.Product
	err = cursor.All(context.Background(), &products)
	return products, err
}

func (r *productRepository) Update(id string, product mongodb.Product) error {
	_, err := r.collection.ReplaceOne(context.Background(), bson.M{"_id": id}, product)
	return err
}

func (r *productRepository) Delete(id string) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

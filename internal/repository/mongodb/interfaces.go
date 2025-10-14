package mongodb

import (
	"api-servers/internal/models/mongodb"
)

type UserRepository interface {
	Create(user mongodb.User) error
	GetByID(id string) (mongodb.User, error)
	GetByEmail(email string) (mongodb.User, error)
	GetAll() ([]mongodb.User, error)
	Update(id string, user mongodb.User) error
	Delete(id string) error
}

type ProductRepository interface {
	Create(product mongodb.Product) error
	GetByID(id string) (mongodb.Product, error)
	GetByCategory(category string) ([]mongodb.Product, error)
	GetAll() ([]mongodb.Product, error)
	Update(id string, product mongodb.Product) error
	Delete(id string) error
}

type OrderRepository interface {
	Create(order mongodb.Order) error
	GetByID(id string) (mongodb.Order, error)
	GetByUserID(user_id string) ([]mongodb.Order, error)
	GetByStatus(status string) ([]mongodb.Order, error)
	GetAll() ([]mongodb.Order, error)
	Update(id string, order mongodb.Order) error
	Delete(id string) error
}

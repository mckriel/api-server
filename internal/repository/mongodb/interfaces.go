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

package mysql

import (
	"api-servers/internal/models/mysql"
)

type customerRepository struct {
	db *Database
}

func NewCustomerRepository(db *Database) CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

func (r *customerRepository) Create(customer mysql.Customer) error {
	var err error
	return err
}

func (r *customerRepository) GetByID(id string) (mysql.Customer, error) {
	var customer mysql.Customer
	var err error
	return customer, err
}

func (r *customerRepository) GetByEmail(email string) (mysql.Customer, error) {
	var customer mysql.Customer
	var err error
	return customer, err
}

func (r *customerRepository) GetByPhone(email string) (mysql.Customer, error) {
	var customer mysql.Customer
	var err error
	return customer, err
}

func (r *customerRepository) GetByName(first_name, last_name string) (mysql.Customer, error) {
	var customer mysql.Customer
	var err error
	return customer, err
}

func (r *customerRepository) GetAll() ([]mysql.Customer, error) {
	var customer []mysql.Customer
	var err error
	return customer, err
}

func (r *customerRepository) Update(id string, customer mysql.Customer) error {
	var err error
	return err
}

func (r *customerRepository) Delete(id string) error {
	var err error
	return err
}

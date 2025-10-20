package mysql

import (
	"api-servers/internal/models/mysql"
	"database/sql"
	"fmt"
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
	query := `INSERT INTO customers (id, first_name, last_name, email, phone, address, city, state, zip_code, date_of_birth, credit_score, created_at, updated_at)
			VALUES (:id, :first_name, :last_name, :email, :phone, :address, :city, :state, :zip_code, :date_of_birth, :credit_score, :created_at, :updated_at)`
	_, err := r.db.Connection.NamedExec(query, customer)

	if err != nil {
		return fmt.Errorf("failed to create customer: %w", err)
	}
	return nil
}

func (r *customerRepository) GetByID(id string) (mysql.Customer, error) {
	var customer mysql.Customer
	err := r.db.Connection.Get(&customer, "SELECT * FROM customers WHERE id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("customer with id %s not found", id)
		}
		return customer, fmt.Errorf("failed to get customer: %w", err)
	}
	return customer, nil
}

func (r *customerRepository) GetByEmail(email string) (mysql.Customer, error) {
	var customer mysql.Customer
	err := r.db.Connection.Get(&customer, "SELECT * FROM customers WHERE email = ?", email)

	if err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("customer with email %s not found", email)
		}
		return customer, fmt.Errorf("failed to get customer: %w", err)
	}
	return customer, nil
}

func (r *customerRepository) GetByPhone(phone string) (mysql.Customer, error) {
	var customer mysql.Customer
	err := r.db.Connection.Get(&customer, "SELECT * FROM customers WHERE phone = ?", phone)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("customer with phone %s not found", phone)
		}
		return customer, fmt.Errorf("failed to get customer: %w", err)
	}
	return customer, nil
}

func (r *customerRepository) GetByName(first_name, last_name string) (mysql.Customer, error) {
	var customer mysql.Customer
	err := r.db.Connection.Get(&customer, "SELECT * FROM customers WHERE first_name = ? AND last_name = ?", first_name, last_name)
	if err != nil {
		if err == sql.ErrNoRows {
			return customer, fmt.Errorf("customer with name %s %s not found", first_name, last_name)
		}
		return customer, fmt.Errorf("failed to get customer: %w", err)
	}
	return customer, nil
}

func (r *customerRepository) GetAll() ([]mysql.Customer, error) {
	var customers []mysql.Customer
	err := r.db.Connection.Select(&customers, "SELECT * FROM customers")
	if err != nil {
		return customers, fmt.Errorf("failed to get customers")
	}
	return customers, nil
}

func (r *customerRepository) Update(id string, customer mysql.Customer) error {
	query := `UPDATE customers SET
                first_name = :first_name,
                last_name = :last_name,
                email = :email,
                phone = :phone,
                address = :address,
                city = :city,
                state = :state,
                zip_code = :zip_code,
                date_of_birth = :date_of_birth,
                credit_score = :credit_score,
                updated_at = :updated_at
                WHERE id = :id`
	customer.ID = id

	result, err := r.db.Connection.NamedExec(query, customer)
	if err != nil {
		return fmt.Errorf("failed to update customer: %w", err)
	}

	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("customer with id %s not found", id)
	}
	return nil
}

func (r *customerRepository) Delete(id string) error {
	query := `DELETE FROM customers WHERE id = ?`

	result, err := r.db.Connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete customer: %w", err)
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected: %w", err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("customer with id %s not found", id)
	}
	return nil
}

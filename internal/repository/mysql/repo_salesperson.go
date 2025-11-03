package mysql

import (
	"api-servers/internal/models/mysql"
	"database/sql"
	"fmt"
)

type salespersonRepository struct {
	db *Database
}

func NewSalespersonRepository(db *Database) SalespersonRepository {
	return &salespersonRepository{
		db: db,
	}
}

func (r *salespersonRepository) Create(salesperson mysql.Salesperson) error {
	query := `INSERT INTO salespersons (id, employee_id, first_name, last_name, email, phone, hire_date, commission, department, status, created_at, updated_at)
			  VALUES (:id, :employee_id, :first_name, :last_name, :email, :phone, :hire_date, :commission, :department, :status, :created_at, :updated_at)`
	_, err := r.db.Connection.NamedExec(query, salesperson)
	if err != nil {
		return fmt.Errorf("failed to create salesperson with employee_id %s: %w", salesperson.Employee_ID, err)
	}
	return nil
}

func (r *salespersonRepository) GetByID(id string) (mysql.Salesperson, error) {
	var salesperson mysql.Salesperson
	err := r.db.Connection.Get(&salesperson, "SELECT * FROM salespersons WHERE id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return salesperson, fmt.Errorf("salesperson with id %s not found", id)
		}
		return salesperson, fmt.Errorf("failed to get salesperson by id %s: %w", id, err)
	}
	return salesperson, nil
}

func (r *salespersonRepository) GetByEmployeeId(employeeId string) (mysql.Salesperson, error) {
	var salesperson mysql.Salesperson
	err := r.db.Connection.Get(&salesperson, "SELECT * FROM salespersons WHERE employee_id = ?", employeeId)

	if err != nil {
		if err == sql.ErrNoRows {
			return salesperson, fmt.Errorf("salesperson with employee_id %s not found", employeeId)
		}
		return salesperson, fmt.Errorf("failed to get salesperson by employee_id %s: %w", employeeId, err)
	}
	return salesperson, nil
}

func (r *salespersonRepository) GetByEmail(email string) (mysql.Salesperson, error) {
	var salesperson mysql.Salesperson
	err := r.db.Connection.Get(&salesperson, "SELECT * FROM salespersons WHERE email = ?", email)

	if err != nil {
		if err == sql.ErrNoRows {
			return salesperson, fmt.Errorf("salesperson with email %s not found", email)
		}
		return salesperson, fmt.Errorf("failed to get salesperson by email %s: %w", email, err)
	}
	return salesperson, nil
}

func (r *salespersonRepository) GetByDepartment(department string) ([]mysql.Salesperson, error) {
	var salespersons []mysql.Salesperson
	err := r.db.Connection.Select(&salespersons, "SELECT * FROM salespersons WHERE department = ?", department)

	if err != nil {
		return salespersons, fmt.Errorf("failed to get salespersons by department %s: %w", department, err)
	}
	return salespersons, nil
}

func (r *salespersonRepository) GetByStatus(status mysql.SalesPersonStatus) ([]mysql.Salesperson, error) {
	var salespersons []mysql.Salesperson
	err := r.db.Connection.Select(&salespersons, "SELECT * FROM salespersons WHERE status = ?", status)

	if err != nil {
		return salespersons, fmt.Errorf("failed to get salespersons by status %s: %w", status, err)
	}
	return salespersons, nil
}

func (r *salespersonRepository) GetAll() ([]mysql.Salesperson, error) {
	var salespersons []mysql.Salesperson
	err := r.db.Connection.Select(&salespersons, "SELECT * FROM salespersons")
	if err != nil {
		return salespersons, fmt.Errorf("failed to get all salespersons: %w", err)
	}
	return salespersons, nil
}

func (r *salespersonRepository) Update(id string, salesperson mysql.Salesperson) error {
	query := `UPDATE salespersons SET
				employee_id = :employee_id,
				first_name = :first_name,
				last_name = :last_name,
				email = :email,
				phone = :phone,
				hire_date = :hire_date,
				commission = :commission,
				department = :department,
				status = :status,
				updated_at = :updated_at
				WHERE id = :id`
	salesperson.ID = id

	result, err := r.db.Connection.NamedExec(query, salesperson)
	if err != nil {
		return fmt.Errorf("failed to update salesperson %s: %w", id, err)
	}

	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for salesperson %s update: %w", id, err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("salesperson with id %s not found for update", id)
	}
	return nil
}

func (r *salespersonRepository) Delete(id string) error {
	query := `DELETE FROM salespersons WHERE id = ?`

	result, err := r.db.Connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete salesperson %s: %w", id, err)
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for salesperson %s deletion: %w", id, err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("salesperson with id %s not found for deletion", id)
	}
	return nil
}
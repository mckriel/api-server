package mysql

import (
	"api-servers/internal/models/mysql"
	"database/sql"
	"fmt"
)

type saleRepository struct {
	db *Database
}

func NewSaleRepository(db *Database) SaleRepository {
	return &saleRepository{
		db: db,
	}
}

func (r *saleRepository) Create(sale mysql.Sale) error {
	query := `INSERT INTO sales (id, vehicle_id, customer_id, salesperson_id, sale_date, sale_price, down_payment, finance_amount, finance_term, interest_rate, payment_method, status, notes, created_at, updated_at)
			  VALUES (:id, :vehicle_id, :customer_id, :salesperson_id, :sale_date, :sale_price, :down_payment, :finance_amount, :finance_term, :interest_rate, :payment_method, :status, :notes, :created_at, :updated_at)`
	_, err := r.db.Connection.NamedExec(query, sale)
	if err != nil {
		return fmt.Errorf("failed to create sale with id %s: %w", sale.ID, err)
	}
	return nil
}

func (r *saleRepository) GetByID(id string) (mysql.Sale, error) {
	var sale mysql.Sale
	err := r.db.Connection.Get(&sale, "SELECT * FROM sales WHERE id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return sale, fmt.Errorf("sale with id %s not found", id)
		}
		return sale, fmt.Errorf("failed to get sale by id %s: %w", id, err)
	}
	return sale, nil
}

func (r *saleRepository) GetByCustomerId(customerId string) ([]mysql.Sale, error) {
	var sales []mysql.Sale
	err := r.db.Connection.Select(&sales, "SELECT * FROM sales WHERE customer_id = ?", customerId)

	if err != nil {
		return sales, fmt.Errorf("failed to get sales by customer_id %s: %w", customerId, err)
	}
	return sales, nil
}

func (r *saleRepository) GetBySalespersonId(salespersonId string) ([]mysql.Sale, error) {
	var sales []mysql.Sale
	err := r.db.Connection.Select(&sales, "SELECT * FROM sales WHERE salesperson_id = ?", salespersonId)

	if err != nil {
		return sales, fmt.Errorf("failed to get sales by salesperson_id %s: %w", salespersonId, err)
	}
	return sales, nil
}

func (r *saleRepository) GetByVehicleId(vehicleId string) (mysql.Sale, error) {
	var sale mysql.Sale
	err := r.db.Connection.Get(&sale, "SELECT * FROM sales WHERE vehicle_id = ?", vehicleId)

	if err != nil {
		if err == sql.ErrNoRows {
			return sale, fmt.Errorf("sale with vehicle_id %s not found", vehicleId)
		}
		return sale, fmt.Errorf("failed to get sale by vehicle_id %s: %w", vehicleId, err)
	}
	return sale, nil
}

func (r *saleRepository) GetByStatus(status mysql.SaleStatus) ([]mysql.Sale, error) {
	var sales []mysql.Sale
	err := r.db.Connection.Select(&sales, "SELECT * FROM sales WHERE status = ?", status)

	if err != nil {
		return sales, fmt.Errorf("failed to get sales by status %s: %w", status, err)
	}
	return sales, nil
}

func (r *saleRepository) GetByPaymentMethod(method mysql.PaymentMethod) ([]mysql.Sale, error) {
	var sales []mysql.Sale
	err := r.db.Connection.Select(&sales, "SELECT * FROM sales WHERE payment_method = ?", method)

	if err != nil {
		return sales, fmt.Errorf("failed to get sales by payment_method %s: %w", method, err)
	}
	return sales, nil
}

func (r *saleRepository) GetByDateRange(startDate, endDate string) ([]mysql.Sale, error) {
	var sales []mysql.Sale
	err := r.db.Connection.Select(&sales, "SELECT * FROM sales WHERE sale_date BETWEEN ? AND ?", startDate, endDate)

	if err != nil {
		return sales, fmt.Errorf("failed to get sales by date range %s to %s: %w", startDate, endDate, err)
	}
	return sales, nil
}

func (r *saleRepository) GetAll() ([]mysql.Sale, error) {
	var sales []mysql.Sale
	err := r.db.Connection.Select(&sales, "SELECT * FROM sales")
	if err != nil {
		return sales, fmt.Errorf("failed to get all sales: %w", err)
	}
	return sales, nil
}

func (r *saleRepository) Update(id string, sale mysql.Sale) error {
	query := `UPDATE sales SET
				vehicle_id = :vehicle_id,
				customer_id = :customer_id,
				salesperson_id = :salesperson_id,
				sale_date = :sale_date,
				sale_price = :sale_price,
				down_payment = :down_payment,
				finance_amount = :finance_amount,
				finance_term = :finance_term,
				interest_rate = :interest_rate,
				payment_method = :payment_method,
				status = :status,
				notes = :notes,
				updated_at = :updated_at
				WHERE id = :id`
	sale.ID = id

	result, err := r.db.Connection.NamedExec(query, sale)
	if err != nil {
		return fmt.Errorf("failed to update sale %s: %w", id, err)
	}

	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for sale %s update: %w", id, err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("sale with id %s not found for update", id)
	}
	return nil
}

func (r *saleRepository) Delete(id string) error {
	query := `DELETE FROM sales WHERE id = ?`

	result, err := r.db.Connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete sale %s: %w", id, err)
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for sale %s deletion: %w", id, err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("sale with id %s not found for deletion", id)
	}
	return nil
}
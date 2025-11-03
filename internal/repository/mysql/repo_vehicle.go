package mysql

import (
	"api-servers/internal/models/mysql"
	"database/sql"
	"fmt"
)

type vehicleRepository struct {
	db *Database
}

func NewVehicleRepository(db *Database) VehicleRepository {
	return &vehicleRepository{
		db: db,
	}
}

func (r *vehicleRepository) Create(vehicle mysql.Vehicle) error {
	query := `INSERT INTO vehicles (id, vin, make, model, year, color, mileage, price, status, engine_type, transmission, fuel_type, created_at, updated_at)
			VALUES (:id, :vin, :make, :model, :year, :color, :mileage, :price, :status, :engine_type, :transmission, :fuel_type, :created_at, :updated_at)`
	_, err := r.db.Connection.NamedExec(query, vehicle)
	if err != nil {
		return fmt.Errorf("failed to create vehicle with VIN %s: %w", vehicle.VIN, err)
	}
	return nil
}

func (r *vehicleRepository) GetByID(id string) (mysql.Vehicle, error) {
	var vehicle mysql.Vehicle
	err := r.db.Connection.Get(&vehicle, "SELECT * FROM vehicles WHERE id = ?", id)

	if err != nil {
		if err == sql.ErrNoRows {
			return vehicle, fmt.Errorf("vehicle with id %s not found", id)
		}
		return vehicle, fmt.Errorf("failed to get vehicle by id %s: %w", id, err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetByVin(vin string) (mysql.Vehicle, error) {
	var vehicle mysql.Vehicle
	err := r.db.Connection.Get(&vehicle, "SELECT * FROM vehicles WHERE vin = ?", vin)

	if err != nil {
		if err == sql.ErrNoRows {
			return vehicle, fmt.Errorf("vehicle with vin %s not found", vin)
		}
		return vehicle, fmt.Errorf("failed to get vehicle by vin %s: %w", vin, err)
	}
	return vehicle, nil

}

func (r *vehicleRepository) GetByMake(make string) ([]mysql.Vehicle, error) {
	var vehicle []mysql.Vehicle
	err := r.db.Connection.Select(&vehicle, "SELECT * FROM vehicles WHERE make = ?", make)

	if err != nil {
		return vehicle, fmt.Errorf("failed to get vehicles by make %s: %w", make, err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetByStatus(status string) ([]mysql.Vehicle, error) {
	var vehicle []mysql.Vehicle
	err := r.db.Connection.Select(&vehicle, "SELECT * FROM vehicles WHERE status = ?", status)

	if err != nil {
		return vehicle, fmt.Errorf("failed to get vehicles by status %s: %w", status, err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetByPriceRange(min_price, max_price float64) ([]mysql.Vehicle, error) {
	var vehicle []mysql.Vehicle
	err := r.db.Connection.Select(&vehicle, "SELECT * FROM vehicles WHERE price BETWEEN ? AND ?", min_price, max_price)

	if err != nil {
		return vehicle, fmt.Errorf("failed to get vehicles by price range $%.2f-$%.2f: %w", min_price, max_price, err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetAll() ([]mysql.Vehicle, error) {
	var vehicles []mysql.Vehicle
	err := r.db.Connection.Select(&vehicles, "SELECT * FROM vehicles")
	if err != nil {
		return vehicles, fmt.Errorf("failed to get all vehicles: %w", err)
	}
	return vehicles, nil
}

func (r *vehicleRepository) Update(id string, vehicle mysql.Vehicle) error {
	query := `UPDATE vehicles SET
				vin = :vin,
				make = :make,
				model = :model,
				year = :year,
				color = :color,
				mileage = :mileage,
				price = :price,
				status = :status,
				engine_type = :engine_type,
				transmission = :transmission,
				fuel_type = :fuel_type,
				updated_at = :updated_at
				WHERE id = :id`
	vehicle.ID = id

	result, err := r.db.Connection.NamedExec(query, vehicle)
	if err != nil {
		return fmt.Errorf("failed to update vehicle %s: %w", id, err)
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for vehicle %s update: %w", id, err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("vehicle with id %s not found for update", id)
	}
	return nil
}

func (r *vehicleRepository) Delete(id string) error {
	query := `DELETE FROM vehicles WHERE id = ?`

	result, err := r.db.Connection.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete vehicle %s: %w", id, err)
	}
	rows_affected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to get rows affected for vehicle %s deletion: %w", id, err)
	}
	if rows_affected == 0 {
		return fmt.Errorf("vehicle with id %s not found for deletion", id)
	}
	return nil
}

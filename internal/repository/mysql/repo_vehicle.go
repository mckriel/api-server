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
		return fmt.Errorf("failed to create vehicle: %w", err)
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
		return vehicle, fmt.Errorf("failed to get vehicle: %w", err)
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
		return vehicle, fmt.Errorf("failed to get vehicle: %w", err)
	}
	return vehicle, nil

}

func (r *vehicleRepository) GetByMake(make string) ([]mysql.Vehicle, error) {
	var vehicle []mysql.Vehicle
	err := r.db.Connection.Get(&vehicle, "SELECT * FROM vehicles WHERE make = ?", make)

	if err != nil {
		if err == sql.ErrNoRows {
			return vehicle, fmt.Errorf("vehicle with vin %s not found", make)
		}
		return vehicle, fmt.Errorf("failed to get vehicle: %w", err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetByStatus(status string) ([]mysql.Vehicle, error) {
	var vehicle []mysql.Vehicle
	err := r.db.Connection.Get(&vehicle, "SELECT * FROM vehicles WHERE status = ?", status)

	if err != nil {
		if err == sql.ErrNoRows {
			return vehicle, fmt.Errorf("vehicle with vin %s not found", status)
		}
		return vehicle, fmt.Errorf("failed to get vehicle: %w", err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetByPriceRange(min_price, max_price float64) ([]mysql.Vehicle, error) {
	var vehicle []mysql.Vehicle
	err := r.db.Connection.Select(&vehicle, "SELECT * FROM vehicles WHERE price BETWEEN ? AND ?", min_price, max_price)

	if err != nil {
		return vehicle, fmt.Errorf("failed to get vehicles: %w", err)
	}
	return vehicle, nil
}

func (r *vehicleRepository) GetAll() ([]mysql.Vehicle, error) {

}

func (r *vehicleRepository) Update(id string, vehicle mysql.Vehicle) error {

}

func (r *vehicleRepository) Delete(id string) error {

}

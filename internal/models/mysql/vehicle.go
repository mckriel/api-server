package mysql

import "time"

type VehicleStatus string

const (
	VehicleStatusAvailable   VehicleStatus = "available"
	VehicleStatusReserved    VehicleStatus = "reserved"
	VehicleStatusSold        VehicleStatus = "sold"
	VehicleStatusMaintenance VehicleStatus = "maintenance"
	VehicleStatusPending     VehicleStatus = "pending"
)

type Vehicle struct {
	ID           string        `json:"id" db:"id"`
	VIN          string        `json:"vin" db:"vin"`
	Make         string        `json:"make" db:"make"`
	Model        string        `json:"model" db:"model"`
	Year         int           `json:"year" db:"year"`
	Color        string        `json:"color" db:"color"`
	Mileage      int           `json:"mileage" db:"mileage"`
	Price        float64       `json:"price" db:"price"`
	Status       VehicleStatus `json:"status" db:"status"`
	Engine_Type  string        `json:"engine_type" db:"engine_type"`
	Transmission string        `json:"transmission" db:"transmission"`
	Fuel_Type    string        `json:"fuel_type" db:"fuel_type"`
	Created_At   time.Time     `json:"created_at" db:"created_at"`
	Updated_At   time.Time     `json:"updated_at" db:"updated_at"`
}

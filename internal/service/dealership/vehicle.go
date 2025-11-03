package dealership

import (
	"api-servers/internal/models/mysql"
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func (s *service) AddVehicleToInventory(ctx context.Context, vehicle VehicleInput) (*mysql.Vehicle, error) {
	newVehicle := mysql.Vehicle{
		ID:           uuid.New().String(),
		VIN:          vehicle.VIN,
		Make:         vehicle.Make,
		Model:        vehicle.Model,
		Year:         vehicle.Year,
		Color:        vehicle.Color,
		Mileage:      vehicle.Mileage,
		Price:        vehicle.Price,
		Status:       "available",
		Engine_Type:  vehicle.EngineType,
		Transmission: vehicle.Transmission,
		Fuel_Type:    vehicle.FuelType,
		Created_At:   time.Now(),
		Updated_At:   time.Now(),
	}

	err := s.vehicle_repo.Create(newVehicle)
	if err != nil {
		return nil, fmt.Errorf("failed to add %s %s to inventory: %w", vehicle.Make, vehicle.Model, err)
	}

	return &newVehicle, nil
}

func (s *service) GetAllVehicles(ctx context.Context) ([]mysql.Vehicle, error) {
	vehicles, err := s.vehicle_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve vehicles: %w", err)
	}
	return vehicles, nil
}

func (s *service) GetVehicleByID(ctx context.Context, vehicleID string) (*mysql.Vehicle, error) {
	vehicle, err := s.vehicle_repo.GetByID(vehicleID)
	if err != nil {
		return nil, fmt.Errorf("vehicle %s not found: %w", vehicleID, err)
	}
	return &vehicle, nil
}

func (s *service) FindVehiclesForCustomers(ctx context.Context, customerID string, preferences VehiclePreferences) ([]mysql.Vehicle, error) {
	allVehicles, err := s.vehicle_repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get vehicles: %w", err)
	}

	var matchingVehicles []mysql.Vehicle
	for _, vehicle := range allVehicles {
		if s.matchesPreferences(vehicle, preferences) {
			matchingVehicles = append(matchingVehicles, vehicle)
		}
	}

	return matchingVehicles, nil
}

func (s *service) ReserveVehicle(ctx context.Context, vehicleID, customerID string) error {
	vehicle, err := s.vehicle_repo.GetByID(vehicleID)
	if err != nil {
		return fmt.Errorf("could not find vehicle %s for reservation: %w", vehicleID, err)
	}

	if vehicle.Status != mysql.VehicleStatusAvailable {
		return fmt.Errorf("vehicle %s is not available for reservation", vehicleID)
	}

	_, err = s.customer_repo.GetByID(customerID)
	if err != nil {
		return fmt.Errorf("customer %s not found for vehicle reservation: %w", customerID, err)
	}

	vehicle.Status = mysql.VehicleStatusReserved
	vehicle.Updated_At = time.Now()

	err = s.vehicle_repo.Update(vehicleID, vehicle)
	if err != nil {
		return fmt.Errorf("failed to reserve vehicle %s for customer %s: %w", vehicleID, customerID, err)
	}

	return nil
}

// vehicle helper functions

func (s *service) matchesPreferences(vehicle mysql.Vehicle, preferences VehiclePreferences) bool {
	if vehicle.Price < preferences.MinPrice || vehicle.Price > preferences.MaxPrice {
		return false
	}

	if (preferences.MinYear > 0 && vehicle.Year < preferences.MinYear) ||
		(preferences.MaxYear > 0 && vehicle.Year > preferences.MaxYear) {
		return false
	}

	if preferences.MaxMileage > 0 && vehicle.Mileage > preferences.MaxMileage {
		return false
	}

	if len(preferences.Makes) > 0 {
		makeSet := make(map[string]bool, len(preferences.Makes))
		for _, make := range preferences.Makes {
			makeSet[make] = true
		}
		if !makeSet[vehicle.Make] {
			return false
		}
	}

	if len(preferences.FuelTypes) > 0 {
		fuelSet := make(map[mysql.FuelType]bool, len(preferences.FuelTypes))
		for _, fuel := range preferences.FuelTypes {
			fuelSet[fuel] = true
		}
		if !fuelSet[vehicle.Fuel_Type] {
			return false
		}
	}

	return true
}

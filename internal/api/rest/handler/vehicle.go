package handler

import (
	"api-servers/internal/service/dealership"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type VehicleHandler struct {
	dealership_service dealership.DealershipService
}

func NewVehicleHandlerService(service dealership.DealershipService) *VehicleHandler {
	return &VehicleHandler{
		dealership_service: service,
	}
}

// GET /vehicles
func (h *VehicleHandler) GetAllVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vehicles, err := h.dealership_service.GetAllVehicles(r.Context())
	if err != nil {
		log.Printf("Error getting all vehicles: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":  "failed to retrieve vehicles",
			"detail": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

// GET /vehicles/{id}
func (h *VehicleHandler) GetVehicleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	vehicle, err := h.dealership_service.GetVehicleByID(r.Context(), vehicleID)
	if err != nil {
		log.Printf("Error getting vehicle %s: %v", vehicleID, err)
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":      "vehicle not found",
			"vehicle_id": vehicleID,
			"detail":     err.Error(),
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicle)
}

// POST /vehicles
func (h *VehicleHandler) CreateVehicle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var vehicleDetails dealership.VehicleInput

	if err := json.NewDecoder(r.Body).Decode(&vehicleDetails); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	vehicle, err := h.dealership_service.AddVehicleToInventory(r.Context(), vehicleDetails)
	if err != nil {
		log.Printf("Error adding vehicle to inventory: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":  "failed to add vehicle",
			"detail": err.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(vehicle)
}

// POST /vehicles/search
func (h *VehicleHandler) SearchVehicles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var searchRequest struct {
		CustomerID  string                        `json:"customer_id"`
		Preferences dealership.VehiclePreferences `json:"preferences"`
	}

	if err := json.NewDecoder(r.Body).Decode(&searchRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	vehicles, err := h.dealership_service.FindVehiclesForCustomers(
		r.Context(),
		searchRequest.CustomerID,
		searchRequest.Preferences,
	)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to search vehicles",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(vehicles)
}

// PUT /vehicles/{id}/reserve
func (h *VehicleHandler) ReserveVehicle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	vehicleID := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	var reservationRequest struct {
		CustomerID string `json:"customer_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&reservationRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	err := h.dealership_service.ReserveVehicle(r.Context(), vehicleID, reservationRequest.CustomerID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to reserve vehicle",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"message":    "vehicle reserved successfully",
		"vehicle_id": vehicleID,
	})
}

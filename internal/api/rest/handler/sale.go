package handler

import (
	"api-servers/internal/service/dealership"
	"encoding/json"
	"net/http"
)

type SaleHandler struct {
	dealership_service dealership.DealershipService
}

func NewSaleHandler(service dealership.DealershipService) *SaleHandler {
	return &SaleHandler{
		dealership_service: service,
	}
}

// POST /sales/start
func (h *SaleHandler) StartSalesProcess(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var startRequest struct {
		CustomerID    string `json:"customer_id"`
		VehicleID     string `json:"vehicle_id"`
		SalespersonID string `json:"salesperson_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&startRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	salesSession, err := h.dealership_service.StartSalesProcess(
		r.Context(),
		startRequest.CustomerID,
		startRequest.VehicleID,
		startRequest.SalespersonID,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to start sales process",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(salesSession)
}

// POST /sales/financing
func (h *SaleHandler) CalculateFinancing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var financingRequest struct {
		VehicleID   string  `json:"vehicle_id"`
		DownPayment float64 `json:"down_payment"`
		CustomerID  string  `json:"customer_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&financingRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	financingOptions, err := h.dealership_service.CalculateFinancingOperations(
		r.Context(),
		financingRequest.VehicleID,
		financingRequest.DownPayment,
		financingRequest.CustomerID,
	)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to calculate financing options",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(financingOptions)
}

// POST /sales/complete
func (h *SaleHandler) ProcessVehicleSale(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var saleRequest dealership.SaleRequest

	if err := json.NewDecoder(r.Body).Decode(&saleRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	saleResult, err := h.dealership_service.ProcessVehicleSale(r.Context(), saleRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to process vehicle sale",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(saleResult)
}

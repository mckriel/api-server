package handler

import (
	"api-servers/internal/service/dealership"
	"encoding/json"
	"net/http"
)

type ReportingHandler struct {
	dealership_service dealership.DealershipService
}

func NewReportingHandler(service dealership.DealershipService) *ReportingHandler {
	return &ReportingHandler{
		dealership_service: service,
	}
}

// GET /reports/sales
func (h *ReportingHandler) GenerateSalesReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reportRequest dealership.ReportPeriod

	if err := json.NewDecoder(r.Body).Decode(&reportRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	salesReport, err := h.dealership_service.GenerateSalesReport(r.Context(), reportRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to generate sales report",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(salesReport)
}

// GET /reports/performance
func (h *ReportingHandler) GetTopPerformers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var reportRequest dealership.ReportPeriod

	if err := json.NewDecoder(r.Body).Decode(&reportRequest); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	performanceReport, err := h.dealership_service.GetTopPerformers(r.Context(), reportRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to get performance report",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(performanceReport)
}

// GET /reports/inventory
func (h *ReportingHandler) GetInventoryReport(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	inventoryReport, err := h.dealership_service.GetInventoryReport(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to get inventory report",
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(inventoryReport)
}

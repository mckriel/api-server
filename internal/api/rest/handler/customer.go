package handler

import (
	"api-servers/internal/service/dealership"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	dealership_service dealership.DealershipService
}

func NewCustomerHandlerService(service dealership.DealershipService) *CustomerHandler {
	return &CustomerHandler{
		dealership_service: service,
	}
}

// GET /customers
func (h *CustomerHandler) GetAllCustomers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"error": "GetAllCustomers not implemented yet",
	})
}

// GET /customers/{id}
func (h *CustomerHandler) GetAllCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotImplemented)
	json.NewEncoder(w).Encode(map[string]string{
		"error":       "GetCustomerByID not implemented yet",
		"customer_id": customerID,
	})
}

// POST /customers

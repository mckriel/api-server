package handler

import (
	"api-servers/internal/api/rest/middleware"
	"api-servers/internal/service/dealership"
	"encoding/json"
	"log"
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

	customers, err := h.dealership_service.GetAllCustomers(r.Context())

	if err != nil {
		log.Printf("Error getting customers: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error":  "failed to retrieve customers",
			"detail": err.Error(),
		})
		return
	}

	version := middleware.GetVersionFromContext(r.Context())

	if version == middleware.Version20241001 {
		// Remove credit_score for older version
		oldCustomers := make([]map[string]interface{}, len(customers))
		for i, customer := range customers {
			oldCustomers[i] = map[string]interface{}{
				"id":            customer.ID,
				"first_name":    customer.First_Name,
				"last_name":     customer.Last_Name,
				"email":         customer.Email,
				"phone":         customer.Phone,
				"address":       customer.Address,
				"city":          customer.City,
				"state":         customer.State,
				"zip_code":      customer.Zip_Code,
				"date_of_birth": customer.Date_Of_Birth,
				"created_at":    customer.Created_At,
				"updated_at":    customer.Updated_At,
			}
		}
		json.NewEncoder(w).Encode(oldCustomers)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customers)
}

// GET /customers/{id}
func (h *CustomerHandler) GetCustomerByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	w.Header().Set("Content-Type", "application/json")
	customerProfile, err := h.dealership_service.GetCustomerProfile(r.Context(), customerID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":       "customer not found",
			"customer_id": customerID,
		})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(customerProfile)
}

// POST /customers
func (h *CustomerHandler) CreateCustomer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var application dealership.CustomerApplication

	if err := json.NewDecoder(r.Body).Decode(&application); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "invalid request body",
		})
		return
	}

	customer, err := h.dealership_service.RegisterNewCustomer(r.Context(), application)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{
			"error": "failed to create customer",
		})
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(customer)
}

// POST /customers/{id}/credit-application
func (h *CustomerHandler) ProcessCreditApplication(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerID := vars["id"]

	w.Header().Set("Content-Type", "application/json")

	creditDecision, err := h.dealership_service.ProcessCreditApplication(r.Context(), customerID)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{
			"error":       "credit processing failed",
			"customer_id": customerID,
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(creditDecision)
}

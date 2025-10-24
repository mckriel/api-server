package rest

import (
	"api-servers/internal/api/rest/handler"
	"api-servers/internal/service/dealership"

	"github.com/gorilla/mux"
)

func SetupRouter(dealershipService dealership.DealershipService) *mux.Router {
	router := mux.NewRouter()

	customerHandler := handler.NewCustomerHandlerService(dealershipService)
	vehicleHandler := handler.NewVehicleHandlerService(dealershipService)
	salesHandler := handler.NewSaleHandler(dealershipService)
	reportingHandler := handler.NewReportingHandler(dealershipService)

	// customer
	router.HandleFunc("/customers", customerHandler.GetAllCustomers).Methods("GET")
	router.HandleFunc("/customers/{id}", customerHandler.GetCustomerByID).Methods("GET")
	router.HandleFunc("/customers", customerHandler.CreateCustomer).Methods("POST")
	router.HandleFunc("/customers/{id}/credit-application", customerHandler.ProcessCreditApplication).Methods("POST")

	// vehicle
	router.HandleFunc("/vehicles", vehicleHandler.GetAllVehicles).Methods("GET")
	router.HandleFunc("/vehicles/{id}", vehicleHandler.GetVehicleByID).Methods("GET")
	router.HandleFunc("/vehicles", vehicleHandler.CreateVehicle).Methods("POST")
	router.HandleFunc("/vehicles/search", vehicleHandler.SearchVehicles).Methods("POST")
	router.HandleFunc("/vehicles/{id}/reserve", vehicleHandler.ReserveVehicle).Methods("PUT")

	// sales
	router.HandleFunc("/sale/start", salesHandler.StartSalesProcess).Methods("POST")
	router.HandleFunc("/sale/financing", salesHandler.CalculateFinancing).Methods("POST")
	router.HandleFunc("/sale/complete", salesHandler.ProcessVehicleSale).Methods("POST")

	// reporting
	router.HandleFunc("/report/sales", reportingHandler.GenerateSalesReport).Methods("GET")
	router.HandleFunc("/report/performance", reportingHandler.GetTopPerformers).Methods("GET")
	router.HandleFunc("/report/inventory", reportingHandler.GetInventoryReport).Methods("GET")

	return router
}

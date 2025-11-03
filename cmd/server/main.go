package main

import (
	"api-servers/internal/api/rest"
	"api-servers/internal/repository/mysql"
	"api-servers/internal/service/dealership"
	"log"
	"net/http"
)

func main() {
	mysqlDB, err := mysql.GetDatabase()
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}

	customerRepo := mysql.NewCustomerRepository(mysqlDB)
	vehicleRepo := mysql.NewVehicleRepository(mysqlDB)
	salespersonRepo := mysql.NewSalespersonRepository(mysqlDB)
	salesRepo := mysql.NewSaleRepository(mysqlDB)

	dealershipService := dealership.NewService(customerRepo, vehicleRepo, salespersonRepo, salesRepo)

	router := rest.SetupRouter(dealershipService)

	log.Println("Starting API server on http://127.0.0.1:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

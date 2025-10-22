package dealership

import (
	"api-servers/internal/repository/mysql"
)

type service struct {
	customer_repo    mysql.CustomerRepository
	vehicle_repo     mysql.VehicleRepository
	salesperson_repo mysql.SalespersonRepository
	sales_repo       mysql.SaleRepository
}

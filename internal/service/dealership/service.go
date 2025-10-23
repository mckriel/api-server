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

func NewService(
	customer_repo mysql.CustomerRepository,
	vehicle_repo mysql.VehicleRepository,
	salesperson_repo mysql.SalespersonRepository,
	sales_repo mysql.SaleRepository,
) DealershipService {
	return &service{
		customer_repo:    customer_repo,
		vehicle_repo:     vehicle_repo,
		salesperson_repo: salesperson_repo,
		sales_repo:       sales_repo,
	}
}

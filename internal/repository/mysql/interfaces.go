package mysql

import (
	"api-servers/internal/models/mysql"
)

type VehicleRepository interface {
	Create(vehicle mysql.Vehicle) error
	GetByID(id string) (mysql.Vehicle, error)
	GetByVin(vin string) (mysql.Vehicle, error)
	GetByMake(make string) ([]mysql.Vehicle, error)
	GetByStatus(status string) ([]mysql.Vehicle, error)
	GetByPriceRange(minPrice, maxPrice float64) ([]mysql.Vehicle, error)
	GetAll() ([]mysql.Vehicle, error)
	Update(id string, vehicle mysql.Vehicle) error
	Delete(id string) error
}

type CustomerRepository interface {
	Create(customer mysql.Customer) error
	GetByID(id string) (mysql.Customer, error)
	GetByEmail(email string) (mysql.Customer, error)
	GetByPhone(phone string) (mysql.Customer, error)
	GetByName(first_name, last_name string) (mysql.Customer, error)
	GetAll() ([]mysql.Customer, error)
	Update(id string, customer mysql.Customer) error
	Delete(id string) error
}

type SalespersonRepository interface {
	Create(salesperson mysql.Salesperson) error
	GetByID(id string) (mysql.Salesperson, error)
	GetByEmployeeId(employeeId string) (mysql.Salesperson, error)
	GetByEmail(email string) (mysql.Salesperson, error)
	GetByDepartment(department string) ([]mysql.Salesperson, error)
	GetByStatus(status mysql.SalesPersonStatus) ([]mysql.Salesperson, error)
	GetAll() ([]mysql.Salesperson, error)
	Update(id string, salesperson mysql.Salesperson) error
	Delete(id string) error
}

type SaleRepository interface {
	Create(sale mysql.Sale) error
	GetByID(id string) (mysql.Sale, error)
	GetByCustomerId(customerId string) ([]mysql.Sale, error)
	GetBySalespersonId(salespersonId string) ([]mysql.Sale, error)
	GetByVehicleId(vehicleId string) (mysql.Sale, error)
	GetByStatus(status mysql.SaleStatus) ([]mysql.Sale, error)
	GetByPaymentMethod(method mysql.PaymentMethod) ([]mysql.Sale, error)
	GetByDateRange(startDate, endDate string) ([]mysql.Sale, error)
	GetAll() ([]mysql.Sale, error)
	Update(id string, sale mysql.Sale) error
	Delete(id string) error
}

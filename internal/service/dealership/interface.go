package dealership

import (
	"api-servers/internal/models/mysql"
	"context"
	"time"
)

type SalesSessionStatus string

const (
	SalesSessionStatusActive    SalesSessionStatus = "active"
	SalesSessionStatusCompleted SalesSessionStatus = "completed"
	SalesSessionStatusCancelled SalesSessionStatus = "cancelled"
	SalesSessionStatusExpired   SalesSessionStatus = "expired"
)

type DealershipService interface {
	// customer
	RegisterNewCustomer(ctx context.Context, application CustomerApplication) (*mysql.Customer, error)
	ProcessCreditApplication(ctx context.Context, customerID string) (*CreditDecision, error)
	GetCustomerProfile(ctx context.Context, customerID string) (*CustomerProfile, error)

	// vehicle
	AddVehicleToInventory(ctx context.Context, vehicle VehicleInput) (*mysql.Vehicle, error)
	FindVehiclesForCustomers(ctx context.Context, customerID string, preferences VehiclePreferences) ([]mysql.Vehicle, error)
	ReserveVehicle(ctx context.Context, vehicleID, customerID string) error

	// sales
	StartSalesProcess(ctx context.Context, customerID, vehicleID, salespersonID string) (*SalesSession, error)
	CalculateFinancingOperations(ctx context.Context, vehicleID string, downPayment float64, customerID string) (FinancingOptions, error)
	ProcessVehicleSale(ctx context.Context, saleRequest SaleRequest) (*SaleResult, error)

	// reporting
	GenerateSalesReport(ctx context.Context, period ReportPeriod) (*SalesReport, error)
	GetTopPerformers(ctx context.Context, period ReportPeriod) (*PerformanceReport, error)
	GetInventoryReport(ctx context.Context) (*InventoryReport, error)
}

type CustomerApplication struct {
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Email        string    `json:"email"`
	Phone        string    `json:"phone"`
	Address      string    `json:"address"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	ZipCode      string    `json:"zip_code"`
	DateOfBirth  time.Time `json:"date_of_birth"`
	AnnualIncome float64   `json:"annual_income"`
}

type CreditDecision struct {
	CustomerID     string  `json:"customer_id"`
	Approved       bool    `json:"approved"`
	CreditLimit    float64 `json:"credit_limit"`
	InterestRate   float64 `json:"interest_rate"`
	ApprovalReason string  `json:"approval_reason"`
	CreditScore    int     `json:"credit_score"`
}

type CustomerProfile struct {
	Customer        mysql.Customer  `json:"customer"`
	PurchaseHistory []mysql.Sale    `json:"purchase_history"`
	OwnedVehicles   []mysql.Vehicle `json:"owned_vehicles"`
	CreditStatus    CreditDecision  `json:"credit_status"`
	TotalSpent      float64         `json:"total_spent"`
	PreferredBrands []string        `json:"preferred_brands"`
}

type VehicleInput struct {
	VIN          string  `json:"vin"`
	Make         string  `json:"make"`
	Model        string  `json:"model"`
	Year         int     `json:"year"`
	Color        string  `json:"color"`
	Mileage      int     `json:"mileage"`
	Price        float64 `json:"price"`
	EngineType   string  `json:"engine_type"`
	Transmission string  `json:"transmission"`
	FuelType     string  `json:"fuel_type"`
}

type VehiclePreferences struct {
	MaxPrice   float64  `json:"max_price"`
	MinPrice   float64  `json:"min_price"`
	Makes      []string `json:"makes"`
	MaxMileage int      `json:"max_mileage"`
	MaxYear    int      `json:"max_year"`
	MinYear    int      `json:"min_year"`
	FuelTypes  []string `json:"fuel_types"`
}

type SalesSession struct {
	SessionID   string             `json:"session_id"`
	Customer    mysql.Customer     `json:"customer"`
	Vehicle     mysql.Vehicle      `json:"vehicle"`
	Salesperson mysql.Salesperson  `json:"salesperson"`
	StartedAt   time.Time          `json:"started_at"`
	Status      SalesSessionStatus `json:"status"`
}

type SaleRequest struct {
	SessionID      string              `json:"session_id"`
	PaymentMethod  mysql.PaymentMethod `json:"payment_method"`
	DownPayment    float64             `json:"down_payment"`
	TradeInVehicle *string             `json:"trade_in_vehicle"`
	FinancingTerm  int                 `json:"financing_term"`
	Notes          string              `json:"notes"`
}

type SaleResult struct {
	Sale             mysql.Sale       `json:"sale"`
	Contract         SalesContract    `json:"contract"`
	FinancingDetails FinancingDetails `json:"financing_details"`
	Commission       float64          `json:"commission"`
}

type FinancingOptions struct {
	CustomerID string            `json:"customer_id"`
	VehicleID  string            `json:"vehicle_id"`
	Options    []FinancingOption `json:"options"`
}

type FinancingOption struct {
	TermMonths     int     `json:"term_months"`
	InterestRate   float64 `json:"interest_rate"`
	MonthlyPayment float64 `json:"monthly_payment"`
	TotalCost      float64 `json:"total_cost"`
}

type SalesContract struct {
	ContractID  string    `json:"contract_id"`
	SaleID      string    `json:"sale_id"`
	Terms       string    `json:"terms"`
	GeneratedAt time.Time `json:"generated_at"`
}

type FinancingDetails struct {
	LoanAmount     float64 `json:"loan_amount"`
	InterestRate   float64 `json:"interest_rate"`
	MonthlyPayment float64 `json:"monthly_payment"`
	TermMonths     int     `json:"term_months"`
}

type ReportPeriod struct {
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}

type SalesReport struct {
	Period         ReportPeriod       `json:"period"`
	TotalSales     int                `json:"total_sales"`
	TotalRevenue   float64            `json:"total_revenue"`
	AverageRevenue float64            `json:"average_revenue"`
	TopVehicles    []VehicleSalesData `json:"top_vehicles"`
	SalesByStatus  map[string]int     `json:"sales_by_status"`
}

type PerformanceReport struct {
	Period          ReportPeriod             `json:"period"`
	TopSalesperson  mysql.Salesperson        `json:"top_salesperson"`
	SalesPersonData []SalespersonPerformance `json:"salespeople"`
}

type VehicleSalesData struct {
	Vehicle      mysql.Vehicle `json:"vehicle"`
	UnitsSold    int           `json:"units_sold"`
	TotalRevenue float64       `json:"total_revenue"`
}

type SalespersonPerformance struct {
	Salesperson  mysql.Salesperson `json:"salesperson"`
	TotalSales   int               `json:"total_sales"`
	TotalRevenue float64           `json:"total_revenue"`
	Commission   float64           `json:"commission"`
}

type InventoryReport struct {
	TotalVehicles    int                `json:"total_vehicles"`
	ValueByMake      map[string]float64 `json:"value_by_make"`
	VehiclesByStatus map[string]int     `json:"vehicles_by_status"`
	AverageAge       int                `json:"average_age"`
	TopValueVehicles []mysql.Vehicle    `json:"top_value_vehicles"`
}

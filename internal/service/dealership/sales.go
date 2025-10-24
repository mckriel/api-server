package dealership

import (
	"api-servers/internal/models/mysql"
	"context"
	"fmt"
	"math"
	"time"

	"github.com/google/uuid"
)

func (s *service) StartSalesProcess(ctx context.Context, customerID, vehicleID, salespersonID string) (*SalesSession, error) {
	customer, err := s.customer_repo.GetByID(customerID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	vehicle, err := s.vehicle_repo.GetByID(vehicleID)
	if err != nil {
		return nil, fmt.Errorf("vehicle not found: %w", err)
	}

	salesperson, err := s.salesperson_repo.GetByID(salespersonID)
	if err != nil {
		return nil, fmt.Errorf("salesperson not found: %w", err)
	}

	if vehicle.Status != mysql.VehicleStatusAvailable && vehicle.Status != mysql.VehicleStatusReserved {
		return nil, fmt.Errorf("vehicle is not available for sale")
	}

	session := &SalesSession{
		SessionID:   uuid.New().String(),
		Customer:    customer,
		Vehicle:     vehicle,
		Salesperson: salesperson,
		StartedAt:   time.Now(),
		Status:      SalesSessionStatusActive,
	}

	return session, nil
}

func (s *service) CalculateFinancingOperations(ctx context.Context, vehicleID string, downPayment float64, customerID string) (FinancingOptions, error) {
	vehicle, err := s.vehicle_repo.GetByID(vehicleID)
	if err != nil {
		return FinancingOptions{}, fmt.Errorf("vehicle not found: %w", err)
	}

	creditDecision, err := s.ProcessCreditApplication(ctx, customerID)
	if err != nil {
		return FinancingOptions{}, fmt.Errorf("failed to get credit decision: %w", err)
	}

	if !creditDecision.Approved {
		return FinancingOptions{}, fmt.Errorf("customer not approved for financing")
	}

	loanAmount := vehicle.Price - downPayment
	if loanAmount > creditDecision.CreditLimit {
		return FinancingOptions{}, fmt.Errorf("loan amount exceeds credit limit")
	}

	financingOptions := []FinancingOption{
		s.calculateFinancingOption(loanAmount, creditDecision.InterestRate, int(FinancingTerm36Months)),
		s.calculateFinancingOption(loanAmount, creditDecision.InterestRate+0.5, int(FinancingTerm48Months)),
		s.calculateFinancingOption(loanAmount, creditDecision.InterestRate+1.0, int(FinancingTerm60Months)),
		s.calculateFinancingOption(loanAmount, creditDecision.InterestRate+1.5, int(FinancingTerm72Months)),
	}

	return FinancingOptions{
		CustomerID: customerID,
		VehicleID:  vehicleID,
		Options:    financingOptions,
	}, nil
}

func (s *service) ProcessVehicleSale(ctx context.Context, saleRequest SaleRequest) (*SaleResult, error) {
	sale := &mysql.Sale{
		ID:             uuid.New().String(),
		Sale_Date:      time.Now(),
		Sale_Price:     0,
		Down_Payment:   saleRequest.DownPayment,
		Finance_Term:   saleRequest.FinancingTerm,
		Payment_Method: saleRequest.PaymentMethod,
		Status:         mysql.SaleStatusCompleted,
		Notes:          saleRequest.Notes,
		Created_At:     time.Now(),
		Updated_At:     time.Now(),
	}

	err := s.sales_repo.Create(*sale)
	if err != nil {
		return nil, fmt.Errorf("failed to create sale: %w", err)
	}

	contract := SalesContract{
		ContractID:  uuid.New().String(),
		SaleID:      sale.ID,
		Terms:       ContractTermsStandard,
		GeneratedAt: time.Now(),
	}

	financingDetails := FinancingDetails{
		LoanAmount:     saleRequest.DownPayment,
		InterestRate:   5.0,
		MonthlyPayment: 500.0,
		TermMonths:     saleRequest.FinancingTerm,
	}

	return &SaleResult{
		Sale:             *sale,
		Contract:         contract,
		FinancingDetails: financingDetails,
		Commission:       1000.0,
	}, nil
}

func (s *service) calculateFinancingOption(loanAmount float64, annualRate float64, termMonths int) FinancingOption {
	monthlyRate := annualRate / 100 / 12
	monthlyPayment := loanAmount * (monthlyRate * math.Pow(1+monthlyRate, float64(termMonths))) / (math.Pow(1+monthlyRate, float64(termMonths)) - 1)
	totalCost := monthlyPayment * float64(termMonths)

	return FinancingOption{
		TermMonths:     termMonths,
		InterestRate:   annualRate,
		MonthlyPayment: monthlyPayment,
		TotalCost:      totalCost,
	}
}

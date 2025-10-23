package dealership

import (
	"api-servers/internal/models/mysql"
	"context"
	"fmt"
	"hash/fnv"
	"time"

	"github.com/google/uuid"
)

func (s *service) RegisterNewCustomer(ctx context.Context, application CustomerApplication) (*mysql.Customer, error) {
	customer := mysql.Customer{
		ID:            uuid.New().String(),
		First_Name:    application.FirstName,
		Last_Name:     application.LastName,
		Email:         application.Email,
		Phone:         application.Phone,
		Address:       application.Address,
		City:          application.City,
		State:         application.State,
		Zip_Code:      application.ZipCode,
		Date_Of_Birth: &application.DateOfBirth,
		Credit_Score:  0,
		Created_At:    time.Now(),
		Updated_At:    time.Now(),
	}

	err := s.customer_repo.Create(customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

func (s *service) ProcessCreditApplication(ctx context.Context, customerID string) (*CreditDecision, error) {
	customer, err := s.customer_repo.GetByID(customerID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	creditScore := s.calculateCreditScore(customer)
	approved := creditScore >= 650
	creditLimit := float64(0)
	interestRate := 15.0
	approvalReason := "credit score too low"

	if approved {
		creditLimit = float64(creditScore)
		if creditScore >= 750 {
			interestRate = 3.5
			approvalReason = "excellent credit score"
		} else if creditScore >= 700 {
			interestRate = 5.9
			approvalReason = "good credit score"
		} else {
			interestRate = 8.9
			approvalReason = "fair credit score"
		}
	}

	return &CreditDecision{
		CustomerID:     customerID,
		Approved:       approved,
		CreditLimit:    creditLimit,
		InterestRate:   interestRate,
		ApprovalReason: approvalReason,
		CreditScore:    creditScore,
	}, nil
}

func (s *service) GetCustomerProfile(ctx context.Context, customerID string) (*CustomerProfile, error) {
	customer, err := s.customer_repo.GetByID(customerID)
	if err != nil {
		return nil, fmt.Errorf("customer not found: %w", err)
	}

	var purchaseHistory []mysql.Sale
	var ownedVehicles []mysql.Vehicle
	var preferredBrands []string
	totalSpent := float64(0)

	creditDecision, err := s.ProcessCreditApplication(ctx, customerID)
	if err != nil {
		return nil, fmt.Errorf("failed to get credit status: %w", err)
	}

	return &CustomerProfile{
		Customer:        customer,
		PurchaseHistory: purchaseHistory,
		OwnedVehicles:   ownedVehicles,
		CreditStatus:    *creditDecision,
		TotalSpent:      totalSpent,
		PreferredBrands: preferredBrands,
	}, nil
}

// customer helper functions

func (s *service) calculateCreditScore(customer mysql.Customer) int {
	score := 300
	if customer.Credit_Score > 0 {
		return customer.Credit_Score
	}

	if customer.Date_Of_Birth != nil {
		age := time.Now().Year() - customer.Date_Of_Birth.Year()
		if age >= 25 {
			score += 100
		}
		if age >= 35 {
			score += 50
		}
	}

	score += hashString(customer.ID)
	if score > 850 {
		return 850
	}

	return score
}

func hashString(score string) int {
	h := fnv.New32a()
	h.Write([]byte(score))
	return int(h.Sum32() % 200)
}

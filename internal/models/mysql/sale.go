package mysql

import "time"

type PaymentMethod string

const (
	PaymentMethodCash    PaymentMethod = "cash"
	PaymentMethodFinance PaymentMethod = "finance"
	PaymentMethodLease   PaymentMethod = "lease"
)

type SaleStatus string

const (
	SaleStatusPending   SaleStatus = "pending"
	SaleStatusCompleted SaleStatus = "completed"
	SaleStatusCancelled SaleStatus = "cancelled"
)

type Sale struct {
	ID             string        `json:"id" db:"id"`
	Vehicle_ID     string        `json:"vehicle_id" db:"vehicle_id"`
	Customer_ID    string        `json:"customer_id" db:"customer_id"`
	Salesperson_ID string        `json:"salesperson_id" db:"salesperson_id"`
	Sale_Date      time.Time     `json:"sale_date" db:"sale_date"`
	Sale_Price     float64       `json:"sale_price" db:"sale_price"`
	Down_Payment   float64       `json:"down_payment" db:"down_payment"`
	Finance_Amount float64       `json:"finance_amount" db:"finance_amount"`
	Finance_Term   int           `json:"finance_term" db:"finance_term"`
	Interest_Rate  float64       `json:"interest_rate" db:"interest_rate"`
	Payment_Method PaymentMethod `json:"payment_method" db:"payment_method"`
	Status         SaleStatus    `json:"status" db:"status"`
	Notes          string        `json:"notes" db:"notes"`
	Created_At     time.Time     `json:"created_at" db:"created_at"`
	Updated_At     time.Time     `json:"updated_at" db:"updated_at"`
}

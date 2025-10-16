package mysql

import "time"

type SalesPersonStatus string

const (
	SalesPersonStatusActive     SalesPersonStatus = "active"
	SalesPersonStatusInactive   SalesPersonStatus = "inactive"
	SalesPersonStatusTerminated SalesPersonStatus = "terminated"
)

type Salesperson struct {
	ID          string            `json:"id" db:"id"`
	Employee_ID string            `json:"employee_id" db:"employee_id"`
	First_Name  string            `json:"first_name" db:"first_name"`
	Last_Name   string            `json:"last_name" db:"last_name"`
	Email       string            `json:"email" db:"email"`
	Phone       string            `json:"phone" db:"phone"`
	Hire_Date   time.Time         `json:"hire_date" db:"hire_date"`
	Commission  float64           `json:"commission" db:"commission"`
	Department  string            `json:"department" db:"department"`
	Status      SalesPersonStatus `json:"status" db:"status"`
	Created_At  time.Time         `json:"created_at" db:"created_at"`
	Updated_At  time.Time         `json:"updated_at" db:"updated_at"`
}

package mysql

import "time"

type Customer struct {
	ID            string     `json:"id" db:"id"`
	First_Name    string     `json:"first_name" db:"first_name"`
	Last_Name     string     `json:"last_name" db:"last_name"`
	Email         string     `json:"email" db:"email"`
	Phone         string     `json:"phone" db:"phone"`
	Address       string     `json:"address" db:"address"`
	City          string     `json:"city" db:"city"`
	State         string     `json:"state" db:"state"`
	Zip_Code      string     `json:"zip_code" db:"zip_code"`
	Date_Of_Birth *time.Time `json:"date_of_birth" db:"date_of_birth"`
	Credit_Score  int        `json:"credit_score" db:"credit_score"`
	Created_At    time.Time  `json:"created_at" db:"created_at"`
	Updated_At    time.Time  `json:"updated_at" db:"updated_at"`
}

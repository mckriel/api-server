package main

import (
	"api-servers/internal"
	"api-servers/internal/models/mysql"
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

func seed_mysql() error {
	db, err := internal.GetMySQLConnection()
	if err != nil {
		return err
	}
	defer db.Close()

	clearQueries := []string{
		"DELETE FROM sales",
		"DELETE FROM vehicles",
		"DELETE FROM customers",
		"DELETE FROM salespersons",
	}

	log.Println("clearing existing data...")
	for _, query := range clearQueries {
		_, err := db.Exec(query)
		if err != nil {
			return err
		}
	}

	vehicles := create_sample_vehicles()
	customers := create_sample_customers()
	salespersons := create_sample_salespersons()
	sales := create_sample_sales(vehicles, customers, salespersons)

	err = seed_vehicles(db, vehicles)
	if err != nil {
		return err
	}

	err = seed_customers(db, customers)
	if err != nil {
		return err
	}

	err = seed_salespersons(db, salespersons)
	if err != nil {
		return err
	}

	err = seed_sales(db, sales)
	if err != nil {
		return err
	}

	log.Printf("created %d vehicles, %d customers, %d salespersons, %d sales",
		len(vehicles), len(customers), len(salespersons), len(sales))
	return nil
}

func create_sample_vehicles() []mysql.Vehicle {
	now := time.Now()

	return []mysql.Vehicle{
		{
			ID:           uuid.New().String(),
			VIN:          "1HGCM82633A123456",
			Make:         "Honda",
			Model:        "Accord",
			Year:         2022,
			Color:        "Silver",
			Mileage:      15000,
			Price:        28500.00,
			Status:       "available",
			Engine_Type:  "2.0L 4-Cylinder",
			Transmission: "CVT",
			Fuel_Type:    "Gasoline",
			Created_At:   now,
			Updated_At:   now,
		},
		{
			ID:           uuid.New().String(),
			VIN:          "2FMDK3GC4DBA12345",
			Make:         "Ford",
			Model:        "Explorer",
			Year:         2023,
			Color:        "Blue",
			Mileage:      8500,
			Price:        42000.00,
			Status:       "available",
			Engine_Type:  "2.3L Turbo 4-Cylinder",
			Transmission: "10-Speed Automatic",
			Fuel_Type:    "Gasoline",
			Created_At:   now,
			Updated_At:   now,
		},
		{
			ID:           uuid.New().String(),
			VIN:          "5NPE34AF5JH123456",
			Make:         "Hyundai",
			Model:        "Sonata",
			Year:         2021,
			Color:        "White",
			Mileage:      22000,
			Price:        24500.00,
			Status:       "sold",
			Engine_Type:  "2.5L 4-Cylinder",
			Transmission: "8-Speed Automatic",
			Fuel_Type:    "Gasoline",
			Created_At:   now,
			Updated_At:   now,
		},
		{
			ID:           uuid.New().String(),
			VIN:          "1G1BC5SM5J7123456",
			Make:         "Chevrolet",
			Model:        "Camaro",
			Year:         2024,
			Color:        "Red",
			Mileage:      2500,
			Price:        35000.00,
			Status:       "reserved",
			Engine_Type:  "3.6L V6",
			Transmission: "8-Speed Automatic",
			Fuel_Type:    "Gasoline",
			Created_At:   now,
			Updated_At:   now,
		},
		{
			ID:           uuid.New().String(),
			VIN:          "3VWD17AJ9EM123456",
			Make:         "Volkswagen",
			Model:        "Jetta",
			Year:         2023,
			Color:        "Black",
			Mileage:      12000,
			Price:        26800.00,
			Status:       "available",
			Engine_Type:  "1.4L Turbo 4-Cylinder",
			Transmission: "8-Speed Automatic",
			Fuel_Type:    "Gasoline",
			Created_At:   now,
			Updated_At:   now,
		},
	}
}

func create_sample_customers() []mysql.Customer {
	now := time.Now()
	birth_date_1985 := time.Date(1985, 6, 15, 0, 0, 0, 0, time.UTC)
	birth_date_1990 := time.Date(1990, 3, 22, 0, 0, 0, 0, time.UTC)
	birth_date_1978 := time.Date(1978, 12, 5, 0, 0, 0, 0, time.UTC)

	return []mysql.Customer{
		{
			ID:            uuid.New().String(),
			First_Name:    "Michael",
			Last_Name:     "Johnson",
			Email:         "michael.johnson@email.com",
			Phone:         "555-0101",
			Address:       "789 Maple Street",
			City:          "Austin",
			State:         "TX",
			Zip_Code:      "73301",
			Date_Of_Birth: &birth_date_1985,
			Credit_Score:  720,
			Created_At:    now,
			Updated_At:    now,
		},
		{
			ID:            uuid.New().String(),
			First_Name:    "Sarah",
			Last_Name:     "Williams",
			Email:         "sarah.williams@email.com",
			Phone:         "555-0202",
			Address:       "456 Oak Avenue",
			City:          "Dallas",
			State:         "TX",
			Zip_Code:      "75201",
			Date_Of_Birth: &birth_date_1990,
			Credit_Score:  680,
			Created_At:    now,
			Updated_At:    now,
		},
		{
			ID:            uuid.New().String(),
			First_Name:    "Robert",
			Last_Name:     "Davis",
			Email:         "robert.davis@email.com",
			Phone:         "555-0303",
			Address:       "123 Pine Boulevard",
			City:          "Houston",
			State:         "TX",
			Zip_Code:      "77001",
			Date_Of_Birth: &birth_date_1978,
			Credit_Score:  750,
			Created_At:    now,
			Updated_At:    now,
		},
	}
}

func create_sample_salespersons() []mysql.Salesperson {
	now := time.Now()
	hire_date_2020 := time.Date(2020, 1, 15, 0, 0, 0, 0, time.UTC)
	hire_date_2019 := time.Date(2019, 8, 22, 0, 0, 0, 0, time.UTC)

	return []mysql.Salesperson{
		{
			ID:          uuid.New().String(),
			Employee_ID: "EMP001",
			First_Name:  "Jennifer",
			Last_Name:   "Smith",
			Email:       "jennifer.smith@dealership.com",
			Phone:       "555-1001",
			Hire_Date:   hire_date_2020,
			Commission:  0.05,
			Department:  "New Cars",
			Status:      mysql.SalesPersonStatusActive,
			Created_At:  now,
			Updated_At:  now,
		},
		{
			ID:          uuid.New().String(),
			Employee_ID: "EMP002",
			First_Name:  "David",
			Last_Name:   "Brown",
			Email:       "david.brown@dealership.com",
			Phone:       "555-1002",
			Hire_Date:   hire_date_2019,
			Commission:  0.06,
			Department:  "Used Cars",
			Status:      mysql.SalesPersonStatusActive,
			Created_At:  now,
			Updated_At:  now,
		},
	}
}

func create_sample_sales(vehicles []mysql.Vehicle, customers []mysql.Customer, salespersons []mysql.Salesperson) []mysql.Sale {
	now := time.Now()
	sale_date_1 := time.Date(2024, 10, 1, 14, 30, 0, 0, time.UTC)
	sale_date_2 := time.Date(2024, 10, 8, 11, 15, 0, 0, time.UTC)

	return []mysql.Sale{
		{
			ID:             uuid.New().String(),
			Vehicle_ID:     vehicles[2].ID, // Sold Hyundai Sonata
			Customer_ID:    customers[0].ID,
			Salesperson_ID: salespersons[0].ID,
			Sale_Date:      sale_date_1,
			Sale_Price:     24500.00,
			Down_Payment:   5000.00,
			Finance_Amount: 19500.00,
			Finance_Term:   60,
			Interest_Rate:  4.5,
			Payment_Method: mysql.PaymentMethodFinance,
			Status:         mysql.SaleStatusCompleted,
			Notes:          "Customer traded in 2018 Toyota Corolla",
			Created_At:     now,
			Updated_At:     now,
		},
		{
			ID:             uuid.New().String(),
			Vehicle_ID:     vehicles[3].ID, // Reserved Chevrolet Camaro
			Customer_ID:    customers[1].ID,
			Salesperson_ID: salespersons[1].ID,
			Sale_Date:      sale_date_2,
			Sale_Price:     35000.00,
			Down_Payment:   10000.00,
			Finance_Amount: 25000.00,
			Finance_Term:   48,
			Interest_Rate:  3.9,
			Payment_Method: mysql.PaymentMethodFinance,
			Status:         mysql.SaleStatusPending,
			Notes:          "Waiting for financing approval",
			Created_At:     now,
			Updated_At:     now,
		},
	}
}

func seed_vehicles(db *sql.DB, vehicles []mysql.Vehicle) error {
	query := `INSERT INTO vehicles (id, vin, make, model, year, color, mileage, price, status, engine_type, transmission, fuel_type, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, vehicle := range vehicles {
		_, err := db.Exec(query, vehicle.ID, vehicle.VIN, vehicle.Make, vehicle.Model, vehicle.Year,
			vehicle.Color, vehicle.Mileage, vehicle.Price, vehicle.Status, vehicle.Engine_Type,
			vehicle.Transmission, vehicle.Fuel_Type, vehicle.Created_At, vehicle.Updated_At)
		if err != nil {
			return err
		}
		log.Printf("created vehicle: %s %s %d (%s)", vehicle.Make, vehicle.Model, vehicle.Year, vehicle.Status)
	}
	return nil
}

func seed_customers(db *sql.DB, customers []mysql.Customer) error {
	query := `INSERT INTO customers (id, first_name, last_name, email, phone, address, city, state, zip_code, date_of_birth, credit_score, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, customer := range customers {
		_, err := db.Exec(query, customer.ID, customer.First_Name, customer.Last_Name, customer.Email,
			customer.Phone, customer.Address, customer.City, customer.State, customer.Zip_Code,
			customer.Date_Of_Birth, customer.Credit_Score, customer.Created_At, customer.Updated_At)
		if err != nil {
			return err
		}
		log.Printf("created customer: %s %s (%s)", customer.First_Name, customer.Last_Name, customer.Email)
	}
	return nil
}

func seed_salespersons(db *sql.DB, salespersons []mysql.Salesperson) error {
	query := `INSERT INTO salespersons (id, employee_id, first_name, last_name, email, phone, hire_date, commission, department, status, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, salesperson := range salespersons {
		_, err := db.Exec(query, salesperson.ID, salesperson.Employee_ID, salesperson.First_Name,
			salesperson.Last_Name, salesperson.Email, salesperson.Phone, salesperson.Hire_Date,
			salesperson.Commission, salesperson.Department, salesperson.Status, salesperson.Created_At, salesperson.Updated_At)
		if err != nil {
			return err
		}
		log.Printf("created salesperson: %s %s (%s)", salesperson.First_Name, salesperson.Last_Name, salesperson.Department)
	}
	return nil
}

func seed_sales(db *sql.DB, sales []mysql.Sale) error {
	query := `INSERT INTO sales (id, vehicle_id, customer_id, salesperson_id, sale_date, sale_price, down_payment, finance_amount, finance_term, interest_rate, payment_method, status, notes, created_at, updated_at) 
			  VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	for _, sale := range sales {
		_, err := db.Exec(query, sale.ID, sale.Vehicle_ID, sale.Customer_ID, sale.Salesperson_ID,
			sale.Sale_Date, sale.Sale_Price, sale.Down_Payment, sale.Finance_Amount, sale.Finance_Term,
			sale.Interest_Rate, sale.Payment_Method, sale.Status, sale.Notes, sale.Created_At, sale.Updated_At)
		if err != nil {
			return err
		}
		log.Printf("created sale: %s ($%.2f, %s)", sale.ID[:8], sale.Sale_Price, sale.Status)
	}
	return nil
}

CREATE TABLE vehicles (
    id VARCHAR(36) PRIMARY KEY,
    vin VARCHAR(17) UNIQUE NOT NULL,
    make VARCHAR(50) NOT NULL,
    model VARCHAR(50) NOT NULL,
    year INT NOT NULL,
    color VARCHAR(30) NOT NULL,
    mileage INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    status ENUM('available', 'sold', 'reserved') DEFAULT 'available',
    engine_type VARCHAR(100),
    transmission VARCHAR(50),
    fuel_type VARCHAR(20),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE customers (
    id VARCHAR(36) PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    address VARCHAR(200),
    city VARCHAR(50),
    state VARCHAR(20),
    zip_code VARCHAR(10),
    date_of_birth DATE,
    credit_score INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE salespersons (
    id VARCHAR(36) PRIMARY KEY,
    employee_id VARCHAR(20) UNIQUE NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    phone VARCHAR(20),
    hire_date DATE NOT NULL,
    commission DECIMAL(4,3) NOT NULL,
    department VARCHAR(50),
    status ENUM('active', 'inactive', 'terminated') DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE sales (
    id VARCHAR(36) PRIMARY KEY,
    vehicle_id VARCHAR(36) NOT NULL,
    customer_id VARCHAR(36) NOT NULL,
    salesperson_id VARCHAR(36) NOT NULL,
    sale_date TIMESTAMP NOT NULL,
    sale_price DECIMAL(10,2) NOT NULL,
    down_payment DECIMAL(10,2) DEFAULT 0,
    finance_amount DECIMAL(10,2) DEFAULT 0,
    finance_term INT DEFAULT 0,
    interest_rate DECIMAL(5,3) DEFAULT 0,
    payment_method ENUM('cash', 'finance', 'lease') NOT NULL,
    status ENUM('pending', 'completed', 'cancelled') DEFAULT 'pending',
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    FOREIGN KEY (vehicle_id) REFERENCES vehicles(id) ON DELETE CASCADE,
    FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE,
    FOREIGN KEY (salesperson_id) REFERENCES salespersons(id) ON DELETE CASCADE
);

CREATE INDEX idx_vehicles_make ON vehicles(make);
CREATE INDEX idx_vehicles_status ON vehicles(status);
CREATE INDEX idx_customers_email ON customers(email);
CREATE INDEX idx_sales_customer ON sales(customer_id);
CREATE INDEX idx_sales_salesperson ON sales(salesperson_id);
CREATE INDEX idx_sales_date ON sales(sale_date);
-- Create customers table
CREATE TABLE customers (
    customer_id SERIAL PRIMARY KEY,
    nik VARCHAR(20) UNIQUE NOT NULL,
    full_name VARCHAR(100) NOT NULL,
    legal_name VARCHAR(100) NOT NULL,
    place_of_birth VARCHAR(50),
    date_of_birth DATE,
    salary NUMERIC(10, 2)
);

-- Create transactions table
CREATE TABLE transactions (
    transaction_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    contract_number VARCHAR(50) UNIQUE NOT NULL,
    otr NUMERIC(10, 2),
    admin_fee NUMERIC(10, 2),
    installment_amount NUMERIC(10, 2),
    interest_amount NUMERIC(10, 2),
    asset_name VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id) ON DELETE CASCADE
);

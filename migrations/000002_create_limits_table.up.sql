-- Create limits table
CREATE TABLE limits (
    limit_id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    tenor INT NOT NULL, -- E.g., 1 month, 3 months, etc.
    limit_amount NUMERIC(10, 2),
    FOREIGN KEY (customer_id) REFERENCES customers (customer_id) ON DELETE CASCADE
);

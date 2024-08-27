CREATE TABLE IF NOT EXISTS product_stock (
    id serial PRIMARY KEY,
    c_code VARCHAR(255) NOT NULL,
    date DATE NOT NULL,
    transaction_type INT NOT NULL,
    transaction_id VARCHAR(255) NOT NULL,
    quantity INT NOT NULL,
    total_sum DECIMAL(10, 2) NOT NULL CHECK (total_sum >= 0),
    transaction_status INT NOT NULL
);
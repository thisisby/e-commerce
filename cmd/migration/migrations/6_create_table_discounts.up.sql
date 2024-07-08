CREATE TABLE IF NOT EXISTS discounts
(
    id         SERIAL PRIMARY KEY,
    product_id INT            NOT NULL REFERENCES products (id),
    discount   DECIMAL(10, 2) NOT NULL CHECK (discount >= 0 AND discount <= 100),
    start_date TIMESTAMP      NOT NULL,
    end_date   TIMESTAMP      NOT NULL CHECK (end_date > start_date),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);
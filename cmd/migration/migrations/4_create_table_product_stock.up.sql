CREATE TABLE IF NOT EXISTS product_stock
(
    transaction_id     VARCHAR(255) PRIMARY KEY UNIQUE,
    customer_id       INT NOT NULL REFERENCES users(id),
    date              DATE NOT NULL,
    active           BOOLEAN NOT NULL,
    order_id INT NULL REFERENCES orders(id)
);
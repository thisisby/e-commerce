CREATE TABLE IF NOT EXISTS product_stock_item(
    product_code     VARCHAR(255) NOT NULL REFERENCES products(c_code),
    transaction_id   VARCHAR(255) NOT NULL REFERENCES product_stock(transaction_id),
    quantity         INTEGER NOT NULL,
    amount           DECIMAL(10, 2) NOT NULL,
    transaction_type INTEGER NOT NULL
);
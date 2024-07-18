CREATE TABLE IF NOT EXISTS order_details
(
    id         SERIAL PRIMARY KEY,
    order_id   int            NOT NULL references orders (id) ON DELETE CASCADE,
    product_id int            NOT NULL references products (id) ON DELETE CASCADE,
    quantity   int            NOT NULL,
    price      DECIMAL(10, 2) NOT NULL,
    sub_total  DECIMAL(10, 2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);
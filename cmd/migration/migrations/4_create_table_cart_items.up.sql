CREATE TABLE IF NOT EXISTS cart_items
(
    id         SERIAL PRIMARY KEY,
    user_id    int NOT NULL REFERENCES users(id),
    product_id INT NOT NULL REFERENCES products(id),
    quantity   INT NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);
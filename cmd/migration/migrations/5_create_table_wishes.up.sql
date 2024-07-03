CREATE TABLE IF NOT EXISTS wishes
(
    id         SERIAL PRIMARY KEY,
    user_id    int NOT NULL REFERENCES users(id),
    product_id INT NOT NULL REFERENCES products(id),

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
    FOREIGN KEY (product_id) REFERENCES products (id) ON DELETE CASCADE
);
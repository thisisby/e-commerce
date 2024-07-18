CREATE TABLE IF NOT EXISTS orders
(
    id               SERIAL PRIMARY KEY,
    user_id          int            NOT NULL references users (id) ON DELETE CASCADE,
    total_price      DECIMAL(10, 2) NOT NULL,
    discounted_price DECIMAL(10, 2) NOT NULL,
    status           VARCHAR(255)   NOT NULL DEFAULT 'pending',
    street           VARCHAR(255)   NOT NULL,
    region           VARCHAR(255)   NOT NULL,
    apartment        VARCHAR(255)   NOT NULL,
    created_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);
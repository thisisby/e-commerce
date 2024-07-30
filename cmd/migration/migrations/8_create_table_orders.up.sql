CREATE TABLE IF NOT EXISTS orders
(
    id               SERIAL PRIMARY KEY,
    user_id          int            NOT NULL references users (id) ON DELETE CASCADE,
    total_price      DECIMAL(10, 2) NOT NULL,
    discounted_price DECIMAL(10, 2) NOT NULL,
    status           VARCHAR(255)   NOT NULL DEFAULT 'pending',
    email            VARCHAR(255)   NULL,
    street           VARCHAR(255)   NULL,
    street_num       VARCHAR(255)   NULL,
    region           VARCHAR(255)   NULL,
    apartment        VARCHAR(255)   NULL,
    delivery_method  VARCHAR(255)   NULL,
    city_id          int            NOT NULL references cities (id) ON DELETE CASCADE,
    created_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP,
    updated_at       TIMESTAMP               DEFAULT CURRENT_TIMESTAMP
);
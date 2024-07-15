CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(100)   NOT NULL,
    description TEXT,
    price       DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    image       TEXT           NOT NULL,
    images      TEXT[]         NULL,

    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
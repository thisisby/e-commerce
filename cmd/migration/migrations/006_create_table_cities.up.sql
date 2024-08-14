CREATE TABLE IF NOT EXISTS cities(
    id serial PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    delivery_duration_days INT NOT NULL
);
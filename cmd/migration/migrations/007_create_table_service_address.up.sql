CREATE TABLE IF NOT EXISTS service_addresses (
    id SERIAL PRIMARY KEY,
    city_id INT NOT NULL REFERENCES cities(id),
    address VARCHAR(255) NOT NULL
);
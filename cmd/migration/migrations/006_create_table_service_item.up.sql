CREATE TABLE IF NOT EXISTS service_items (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    duration INT NOT NULL,
    description VARCHAR(255) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    subservice_id INT NOT NULL REFERENCES subservices(id) ON DELETE CASCADE
);
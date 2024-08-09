CREATE TABLE IF NOT EXISTS subservices (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    service_id INT NOT NULL REFERENCES services(id)
);
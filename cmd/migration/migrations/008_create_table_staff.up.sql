CREATE TABLE IF NOT EXISTS staff (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    occupation VARCHAR(255) NOT NULL,
    experience INT NOT NULL,
    avatar VARCHAR(255) NULL,
    service_id INT NOT NULL REFERENCES services(id) ON DELETE CASCADE
);
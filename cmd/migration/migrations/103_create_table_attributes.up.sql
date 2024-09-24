CREATE TABLE IF NOT EXISTS attributes(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    characteristic_id INTEGER NOT NULL REFERENCES characteristics(id)
);
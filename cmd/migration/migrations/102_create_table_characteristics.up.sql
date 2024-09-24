CREATE TABLE IF NOT EXISTS characteristics(
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    subcategory_id INTEGER NOT NULL REFERENCES subcategories(id)
);
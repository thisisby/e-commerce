CREATE TABLE IF NOT EXISTS personal_addresses
(
    id SERIAL PRIMARY KEY,
    street VARCHAR(255) NULL,
    street_num VARCHAR(255) NULL,
    region VARCHAR(255) NULL,
    apartment VARCHAR(255) NULL,

    user_id int DEFAULT 1 REFERENCES users(id) ON DELETE CASCADE,
    city_id int DEFAULT 1 REFERENCES cities(id) ON DELETE CASCADE
);
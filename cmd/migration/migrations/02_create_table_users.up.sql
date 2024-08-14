CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    phone VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    date_of_birth DATE NOT NULL,
    refresh_token VARCHAR(255) NULL,

    email VARCHAR(255) NULL UNIQUE,
    street VARCHAR(255) NULL,
    street_num VARCHAR(255) NULL,
    region VARCHAR(255) NULL,
    apartment VARCHAR(255) NULL,

    city_id int DEFAULT 1 REFERENCES cities(id) ON DELETE CASCADE,
    role_id int DEFAULT 2 REFERENCES roles(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
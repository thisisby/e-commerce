CREATE TABLE IF NOT EXISTS users(
    id serial PRIMARY KEY,
    phone VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(50) NOT NULL,
    date_of_birth DATE NOT NULL,
    refresh_token VARCHAR(255) NULL,

    street VARCHAR(50) NULL,
    region VARCHAR(50) NULL,
    apartment VARCHAR(50) NULL,

    country_id int DEFAULT 1 REFERENCES countries(id),
    role_id int DEFAULT 2 REFERENCES roles(id),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (country_id) REFERENCES countries(id) ON DELETE CASCADE,
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE
);
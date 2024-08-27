CREATE TABLE IF NOT EXISTS products
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100)   NOT NULL,
    description    TEXT,
    ingredients    TEXT,
    c_code         VARCHAR(255)   NOT NULL,
    ed_izm         VARCHAR(255)   NOT NULL,
    article        VARCHAR(255)   NOT NULL,
    price          DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    weight         DECIMAL(10, 2) NULL,
    image          TEXT           NULL,
    images         TEXT[]         NULL,
    subcategory_id INT            NULL REFERENCES subcategories (id),
    brand_id       INT            NULL REFERENCES brands (id),

    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
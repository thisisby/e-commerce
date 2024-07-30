CREATE TABLE IF NOT EXISTS products
(
    id             SERIAL PRIMARY KEY,
    name           VARCHAR(100)   NOT NULL,
    description    TEXT,
    price          DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    image          TEXT           NOT NULL,
    images         TEXT[]         NULL,
    stock          INTEGER        NOT NULL CHECK (stock >= 0),
    subcategory_id INT            NOT NULL REFERENCES subcategories (id),
    brand_id       INT            NOT NULL REFERENCES brands (id),

    created_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at     TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
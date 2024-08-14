CREATE TABLE IF NOT EXISTS appointments (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL REFERENCES users(id),
    staff_id INT NOT NULL REFERENCES staff(id),
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL,
    service_item_id INT NOT NULL REFERENCES service_items(id),
    comments TEXT,
    status VARCHAR(255) NOT NULL
);
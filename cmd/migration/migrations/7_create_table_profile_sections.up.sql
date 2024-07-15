CREATE TABLE IF NOT EXISTS profile_sections
(
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(255) NOT NULL,
    content    TEXT,
    parent_id  int          NULL references profile_sections (id) ON DELETE SET NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (parent_id) REFERENCES profile_sections (id) ON DELETE SET NULL
);
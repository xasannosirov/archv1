CREATE TABLE IF NOT EXISTS menus (
    id SERIAL PRIMARY KEY,
    title JSONB NOT NULL,
    content JSONB,
    is_static BOOLEAN DEFAULT TRUE,
    sort INTEGER NOT NULL,
    parent_id INTEGER,
    status BOOLEAN DEFAULT TRUE,
    slug VARCHAR(100),
    path VARCHAR(100),
    files VARCHAR[],
    created_at TIMESTAMP DEFAULT NOW(),
    created_by INTEGER REFERENCES users(id),
    updated_at TIMESTAMP,
    updated_by INTEGER REFERENCES users(id),
    deleted_at TIMESTAMP,
    deleted_by INTEGER REFERENCES users(id)
);

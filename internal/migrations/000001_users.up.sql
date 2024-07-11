CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(100),
    password TEXT,
    role VARCHAR(10),
    status BOOLEAN NOT NULL DEFAULT TRUE,
    refresh TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    created_by INT,
    updated_at TIMESTAMP,
    updated_by INT,
    deleted_at TIMESTAMP,
    deleted_by INT
);

ALTER TABLE users ADD CONSTRAINT fk_created_by FOREIGN KEY (created_by) REFERENCES users(id);

ALTER TABLE users ADD CONSTRAINT fk_updated_by FOREIGN KEY (updated_by) REFERENCES users(id);

ALTER TABLE users ADD CONSTRAINT fk_deleted_by FOREIGN KEY (deleted_by) REFERENCES users(id);

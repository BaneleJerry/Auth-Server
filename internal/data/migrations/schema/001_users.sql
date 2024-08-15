-- +goose Up
-- +goose StatementBegin
CREATE TABLE
    users (
        id UUID PRIMARY KEY NOT NULL,
        username VARCHAR(255) UNIQUE NOT NULL,
        email VARCHAR(255) UNIQUE NOT NULL,
        password_hash TEXT NOT NULL,
        created_at TIMESTAMPTZ DEFAULT now () NOT NULL,
        updated_at TIMESTAMPTZ DEFAULT now () NOT NULL
    );

CREATE TABLE
    user_profiles (
        user_id UUID PRIMARY KEY REFERENCES users (id) ON DELETE CASCADE,
        first_name VARCHAR(255),
        last_name VARCHAR(255),
        phone_number VARCHAR(20),
        address TEXT,
        created_at TIMESTAMPTZ DEFAULT now () NOT NULL,
        updated_at TIMESTAMPTZ DEFAULT now () NOT NULL
    );

CREATE INDEX idx_users_username ON users (username);

CREATE INDEX idx_users_email ON users (email);

CREATE INDEX idx_user_profiles_user_id ON user_profiles (user_id);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_profiles;

DROP TABLE IF EXISTS users;

-- +goose StatementEnd
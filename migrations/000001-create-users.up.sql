-- +migrate Up

CREATE TABLE IF NOT EXISTS users (
     id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
     name VARCHAR(150) NOT NULL,
     email VARCHAR(255) UNIQUE NOT NULL,
     password TEXT NOT NULL,
     is_owner BOOLEAN NOT NULL DEFAULT FALSE,
     created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
     updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);



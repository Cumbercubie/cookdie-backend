CREATE EXTENSION IF NOT EXISTS "pgcrypto"; -- for gen_random_uuid()

CREATE TABLE restaurants (
    id UUID PRIMARY KEY,
    name TEXT,
    address TEXT,
    city TEXT,
    state VARCHAR(2),
    star INTEGER DEFAULT 0 NOT NULL,
    postal_code VARCHAR(10),
    description TEXT,
    updated_at TIMESTAMP ,
    created_at TIMESTAMP DEFAULT NOW() NOT NULL
);

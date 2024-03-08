DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS friendships;

-- Create users table
CREATE TABLE users (
    id          SERIAL PRIMARY KEY,
    name        varchar    NOT NULL,
    age         INT    NOT NULL,
    created_at  timestamp NOT NULL DEFAULT NOW(),
    updated_at  timestamp NOT NULL DEFAULT NOW()
);

-- Create friendship table
CREATE TABLE friendships (
    id        SERIAL PRIMARY KEY,
    source_id INT    NOT NULL,
    target_id INT    NOT NULL
);

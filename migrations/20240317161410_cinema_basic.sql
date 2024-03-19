-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
  id uuid PRIMARY KEY,
  login varchar(255),
  role varchar(255),
  hashed_password text,
  is_deleted boolean,
  created_at timestamp,
  updated_at timestamp,
  deleted_at timestamp
);

CREATE INDEX user_login ON users(login);

CREATE TABLE IF NOT EXISTS actors (
    id uuid PRIMARY KEY,
    sex char(1),
    bith_date date,
    is_deleted bool,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS movies (
    id uuid PRIMARY KEY,
    title varchar(150),
    description varchar(1000),
    release_date date,
    rating decimal,
    is_deleted boolean,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
);

CREATE TABLE IF NOT EXISTS actors_films (
  actor_id uuid REFERENCES actors (id) ON DELETE CASCADE,
  cloth_id uuid REFERENCES films (id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS actors_films;
DROP TABLE IF EXISTS actors;
DROP TABLE IF EXISTS films;
-- +goose StatementEnd

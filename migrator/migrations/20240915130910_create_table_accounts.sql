-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
CREATE TABLE IF NOT EXISTS accounts(
    id UUID NOT NULL PRIMARY KEY,
    login VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    password VARCHAR(100) NOT NULL,
    phone VARCHAR(50) NOT NULL UNIQUE
);
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE INDEX email_index ON accounts (email);
CREATE INDEX phone_index ON accounts (phone);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS accounts;
DROP EXTENSION IF EXISTS "uuid-ossp";
-- +goose StatementEnd
